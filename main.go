package main

import (
	"fmt"
	"kpay/config"
	"kpay/dataaccessobject"
	"kpay/model"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	configs = config.Config{}
	daos    = dataaccessobject.DataAccessObject{}
)

func init() {
	configs.Read()
	daos.Server = configs.Server
	daos.Database = configs.Database
	dbAccess := daos.ConnectDatabase()
	fmt.Println("Connected Database: ", dbAccess)
}

func main() {
	r := initializeRoutes()
	r.Run(":" + os.Getenv("PORT"))
}

func CreateMerchantEndPoint(c *gin.Context) {

	var register model.Register
	err := c.ShouldBindJSON(&register)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	merchant, err := daos.Register(&register)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, merchant)
}

func AllMerchantEndPoint(c *gin.Context) {
	merchants, err := daos.FindAll()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, merchants)
}

func FindByIdMerchantEndPoint(c *gin.Context) {

	merchant, err := daos.FindById(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, merchant)
}

func UpdateIdMerchantEndPoint(c *gin.Context) {

	merchant, err := daos.FindById(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	errs := c.ShouldBindJSON(&merchant)
	if errs != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := daos.Update(merchant); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{"result": "success"})
}

// func CreateBankAccountOfMerchantEndPoint(c *gin.Context) {

// 	var merchantBankAcc model.BankAccout

// 	merchant, err := daos.FindById(c.Param("id"))
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	errs := c.ShouldBindJSON(&merchant)
// 	if errs != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	merchantBankAcc.ID = bson.NewObjectId()
// 	merchant.BankAccount = append(merchant.BankAccount, merchantBankAcc)

// 	if err := daos.Update(c.Param("id"), merchant); err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, map[string]string{"result": "success"})
// }
