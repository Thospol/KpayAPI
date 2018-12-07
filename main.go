package main

import (
	"fmt"
	"kpay/config"
	"kpay/helper"
	"kpay/logrus"
	"kpay/middleware"
	"kpay/model"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	logr "github.com/sirupsen/logrus"
)

var (
	configs   = config.Config{}
	daos      = middleware.SDO
	formatter = logrus.LogFormat{}
)

func init() {
	configs.Read()
	daos.Server = configs.Server
	daos.Database = configs.Database
	dbAccess := daos.ConnectDatabase()
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	logr.SetFormatter(&formatter)
	log.SetOutput(os.Stderr)
	fmt.Println("Connected Database: ", dbAccess)
}

func main() {
	r := initializeRoutes()
	r.Run(":" + os.Getenv("PORT"))
}

func CreateMerchantEndPoint(c *gin.Context) {

	merchants, err := daos.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var register model.Register
	if err := c.ShouldBindJSON(&register); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	merchantRespone, err := daos.Register(&register, merchants)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	logrus.LogrusfMerchantPOST(merchantRespone)
	c.JSON(http.StatusCreated, map[string]string{"result": "success"})
}

func AllMerchantEndPoint(c *gin.Context) {
	merchants, err := daos.FindAll()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	logrus.LogrusMerchantAll(merchants)
	c.JSON(http.StatusOK, helper.MapData(merchants))
}

func FindByIdMerchantEndPoint(c *gin.Context) {

	merchant, err := daos.FindById(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	logrus.LogrusMerchantByID(merchant)
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

	logrus.LogrusMerchantByID(merchant)
	c.JSON(http.StatusOK, map[string]string{"result": "success"})
}

func FindByIdProductMerchantEndPoint(c *gin.Context) {

	merchant, err := daos.FindById(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	product := daos.FindProductById(&merchant)
	logrus.LogrusProductOfMerchantFindList(product)
	c.JSON(http.StatusOK, helper.MapDataProduct(product))
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
	logrus.LogrusProductPOSTPUT(merchantForResponese)
	c.JSON(http.StatusCreated, helper.MapData(merchantForResponese))
}

func DeleteProductMerchantEndPoint(c *gin.Context) {

	merchant, err := daos.FindById(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	merchantForResponese, err := daos.DeleteProductMerchant(c.Param("product_id"), merchant)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	logrus.LogrusProduct(merchantForResponese)
	c.JSON(http.StatusOK, map[string]string{"result": "success"})
}

func UpdateProductMerchantEndPoint(c *gin.Context) {

	merchant, err := daos.FindById(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	var updateProductMerchant model.UpdateProduct
	if err := c.ShouldBindJSON(&updateProductMerchant); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: require only amount: %s", err),
		})
		return
	}

	merchantResponse, err := daos.UpdateProductMerchant(c.Param("product_id"), &updateProductMerchant, merchant)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	logrus.LogrusProductPOSTPUT(merchantResponse)
	c.JSON(http.StatusOK, map[string]string{"result": "success"})
}

func BuyProductInMerchantEndPoint(c *gin.Context) {
	var requestBuyProduct model.BuyProduct
	if err := c.ShouldBindJSON(&requestBuyProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}
	merchantresponse, err := daos.BuyProductMerchant(&requestBuyProduct)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	logrus.LogrusBuyMerchant(merchantresponse)

	c.JSON(http.StatusCreated, map[string]string{"result": "Buy Success"})
}

func AllReportMerchantEndPoint(c *gin.Context) {

	_, err := daos.FindAll()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	reports, err := daos.FindAllReportMerchant()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	logrus.LogrusReportAll(reports)
	c.JSON(http.StatusOK, helper.MapDataReport(reports))
}

func FindByIdReportMerchantEndPoint(c *gin.Context) {
	merchant, err := daos.FindById(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	report, err := daos.FindReportMerchant(merchant)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	logrus.LogrusReportAll(report)
	c.JSON(http.StatusCreated, helper.MapDataReport(report))
}

func CreateReportMerchantEndPoint(c *gin.Context) {

	merchant, err := daos.FindById(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	var addReport model.AddReport
	if err := c.ShouldBindJSON(&addReport); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}
	merchantReport, err := daos.AddReport(&addReport, merchant)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	logrus.LogrusReportPOST(merchantReport)
	c.JSON(http.StatusCreated, helper.MapData(merchantReport))
}

func CreateUserEndPoint(c *gin.Context) {

	users, err := daos.FindAllUser()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var RequestUser model.User
	if err := c.ShouldBindJSON(&RequestUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	userRep, errs := daos.InsertUser(&RequestUser, users)
	if errs != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	logrus.LogrusUserPOST(userRep)
	c.JSON(http.StatusCreated, map[string]string{"result": "Create Success"})
}

func FindAllUserEndPoint(c *gin.Context) {

	users, err := daos.FindAllUser()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, helper.MapDataUser(users))
}

func FindByIDUserEndPoint(c *gin.Context) {

	user, err := daos.FindByIDUser(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, helper.MapDataUser(user))
}

func CreateBankAccountOfuserEndPoint(c *gin.Context) {
	users, err := daos.FindAllUser()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	user, err := daos.FindByIDUser(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var bankAccountRequest model.UserBankAccount
	if err := c.ShouldBindJSON(&bankAccountRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}
	if bankAccountRequest.BankName == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"Result": "please require BankName"})
		return
	}
	if bankAccountRequest.AccountNumber == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"Result": "please require AccountNumber"})
		return
	}
	if bankAccountRequest.Balance == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"Result": "please require Balance"})
		return
	}
	if err := daos.InsertBankAccountOfUser(&bankAccountRequest,user,users); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,map[string]string{"result":"Insert BankAccount Success"})

}
