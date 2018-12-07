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
	LogrusfMerchantPOST(merchantRespone)
	c.JSON(http.StatusCreated, map[string]string{"result": "success"})
}

func AllMerchantEndPoint(c *gin.Context) {
	merchants, err := daos.FindAll()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	LogrusMerchantAll(merchants)
	c.JSON(http.StatusOK, helper.MapData(merchants))
}

func FindByIdMerchantEndPoint(c *gin.Context) {

	merchant, err := daos.FindById(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	LogrusMerchantByID(merchant)
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

	LogrusMerchantByID(merchant)
	c.JSON(http.StatusOK, map[string]string{"result": "success"})
}

func FindByIdProductMerchantEndPoint(c *gin.Context) {

	merchant, err := daos.FindById(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	product := daos.FindProductById(&merchant)
	LogrusProductOfMerchantFindList(product)
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
	LogrusProductPOSTPUT(merchantForResponese)
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
	LogrusProduct(merchantForResponese)
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
	LogrusProductPOSTPUT(merchantResponse)
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
	LogrusBuyMerchant(merchantresponse)

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
	LogrusReportAll(reports)
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
	LogrusReportAll(report)
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
	LogrusReportPOST(merchantReport)
	c.JSON(http.StatusCreated, helper.MapData(merchantReport))
}

func CreateUserEndPoint(c *gin.Context) {

	var RequestUser model.User
	if err := c.ShouldBindJSON(&RequestUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}
	//ค่อยมาต่อ
}

func LogrusMerchantByID(merchant model.Merchant) {
	logr.WithFields(logr.Fields{
		"\nmerchant.ID":          merchant.ID,
		"\nmerchant.Name":        merchant.Name,
		"\nmerchant.Username":    merchant.Username,
		"\nmerchant.Password":    merchant.Password,
		"\nmerchant.BankAccount": merchant.BankAccount,
		"\nmerchant.Products":    merchant.Products,
		"\nmerchant.Report":      merchant.Report,
	}).Info("Merchant")
}

func LogrusMerchantAll(merchant []model.Merchant) {

	for _, merchantList := range merchant {
		logr.WithFields(logr.Fields{
			"\nmerchant.ID":          merchantList.ID,
			"\nmerchant.Name":        merchantList.Name,
			"\nmerchant.Username":    merchantList.Username,
			"\nmerchant.Password":    merchantList.Password,
			"\nmerchant.BankAccount": merchantList.BankAccount,
			"\nmerchant.Products":    merchantList.Products,
			"\nmerchant.Report":      merchantList.Report,
		}).Info("Merchant")
	}
}

func LogrusProductOfMerchantFindList(product []model.Product) {

	for _, prodcutList := range product {
		logr.WithFields(logr.Fields{
			"\nprodcutList.IDMerchant":   prodcutList.IDMerchant,
			"\nprodcutList.ID":           prodcutList.ID,
			"\nprodcutList.NameProduct":  prodcutList.NameProduct,
			"\nprodcutList.Amount":       prodcutList.Amount,
			"\nprodcutList.Volume":       prodcutList.Volume,
			"\nprodcutList.AmountChange": prodcutList.AmountChange,
		}).Info("Merchant -> Prodcut")
	}
}

func LogrusfMerchantPOST(merchant *model.Merchant) {

	logr.WithFields(logr.Fields{
		"\nmerchant.ID":          merchant.ID,
		"\nmerchant.Name":        merchant.Name,
		"\nmerchant.Username":    merchant.Username,
		"\nmerchant.Password":    merchant.Password,
		"\nmerchant.BankAccount": merchant.BankAccount,
	}).Info("Merchant")
}

func LogrusProduct(merchant model.Merchant) {

	logr.WithFields(logr.Fields{
		"\nmerchant.Products": merchant.Products,
	}).Info("Merchant -> Prodcut")
}

func LogrusProductPOSTPUT(merchant *model.Merchant) {

	logr.WithFields(logr.Fields{
		"\nmerchant.Products": merchant.Products,
	}).Info("Merchant -> Prodcut")
}

func LogrusBuyMerchant(merchant *model.Merchant) {
	logr.WithFields(logr.Fields{
		"\nmerchant.ID":          merchant.ID,
		"\nmerchant.Name":        merchant.Name,
		"\nmerchant.Username":    merchant.Username,
		"\nmerchant.Password":    merchant.Password,
		"\nmerchant.BankAccount": merchant.BankAccount,
		"\nmerchant.Products":    merchant.Products,
		"\nmerchant.Report":      merchant.Report,
	}).Info("Merchant")
}

func LogrusReportAll(report []model.Report) {
	for _, reportList := range report {
		logr.WithFields(logr.Fields{
			"\nreport.Date":          reportList.Date,
			"\nreport.ID": reportList.ID,
			"\nreport.ProductSelling":      reportList.ProductSelling,
			"\nreport.Accumulate":      reportList.Accumulate,
		}).Info("Merchant -> Report")
	}
}

func LogrusReportPOST(merchant *model.Merchant) {
	for _, listMerchantReport := range merchant.Report {
		logr.WithFields(logr.Fields{
			"\nreport.Date":          listMerchantReport.Date,
			"\nreport.ID": listMerchantReport.ID,
			"\nreport.ProductSelling":      listMerchantReport.ProductSelling,
			"\nreport.Accumulate":      listMerchantReport.Accumulate,
		}).Info("Merchant -> Report")
	}
}
