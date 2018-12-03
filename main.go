package main

import (
	"fmt"
	"kpay/config"
	"kpay/dataaccessobject"
	"kpay/helper"
	"kpay/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
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
	r.Run(":3000")
}

func CreateMerchantEndPoint(c *gin.Context) {

	var merchant model.Merchant
	err := c.ShouldBindJSON(&merchant)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	merchant.ID = bson.NewObjectId()
	merchant.Username = helper.RandomUsername()
	merchant.Password = helper.RandomPassword()
	if err := daos.Insert(merchant); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, merchant)
}