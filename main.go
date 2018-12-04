package main

import (
	"fmt"
	"kpay/config"
	"kpay/dataaccessobject"
	"kpay/helper"
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

	if _, err := daos.Register(&register); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]string{"result": "success"})
}

func AllMerchantEndPoint(c *gin.Context) {
	merchants, err := daos.FindAll()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helper.MapData(merchants))
}

func FindByIdMerchantEndPoint(c *gin.Context) {

	merchant, err := daos.FindById(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helper.MapData(merchant))
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

func CreateProductMerchantEndPoint(c *gin.Context) {

	var addproduct model.AddProduct
	err := c.ShouldBindJSON(&addproduct)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}
	merchant, err := daos.FindById(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	merchantForResponese, err := daos.AddProduct(&addproduct, merchant)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helper.MapData(merchantForResponese))
}

func DeleteProductMerchantEndPoint(c *gin.Context) {

	merchant, err := daos.FindById(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if _, err := daos.DeleteProductMerchant(c.Param("product_id"), merchant); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]string{"result": "success"})
}
