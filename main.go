package main

import (
	"fmt"
	"kpay/config"
	"kpay/helper"
	"kpay/middleware"
	"kpay/model"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

var (
	configs = config.Config{}
	daos    = middleware.SDO
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

	if _, err := daos.Register(&register, merchants); err != nil {
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

	if _, err := daos.UpdateProductMerchant(c.Param("product_id"), &updateProductMerchant, merchant); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{"result": "success"})
}

func BuyProductInMerchantEndPoint(c *gin.Context) {

	var buyProduct model.BuyProduct
	var report model.Report
	var productSellingReport model.ProductSellingReport

	merchants, err := daos.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&buyProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if buyProduct.ProductName == "" {
		c.JSON(http.StatusInternalServerError, map[string]string{"result": "please require ProductName"})
		return
	}
	if buyProduct.Volume == 0 {
		c.JSON(http.StatusInternalServerError, map[string]string{"result": "please require Volume"})
		return
	}
	var priceOfProductTotal float64
	var idMerchant bson.ObjectId
	hasProduct := false
	for _, listOfMerchant := range merchants {
		for _, listProductInMerchant := range listOfMerchant.Products {
			if listProductInMerchant.NameProduct == buyProduct.ProductName {
				productSellingReport.ID = bson.NewObjectId()
				productSellingReport.Name = buyProduct.ProductName
				productSellingReport.SellingVolume = buyProduct.Volume
				priceOfProductTotal = listProductInMerchant.Amount * float64(buyProduct.Volume)
				idMerchant = listProductInMerchant.IDMerchant
				hasProduct = true
			}
		}
	}
	if hasProduct == false {
		c.JSON(http.StatusInternalServerError, map[string]string{"result": "don't have Product in Merchant"})
		return
	}

	reports, err := daos.FindAllReport()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	hasMerchant := false
	for _, listReport := range reports {
		if listReport.IDMerchant == idMerchant {
			hasMerchant = true
			if listReport.Date == time.Now().Format("02-01-2006") {
				report, err = daos.FindByIdReport(listReport.ID)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
					return
				}
				report.Accumulate = report.Accumulate + priceOfProductTotal
				report.ProductSelling = append(report.ProductSelling, productSellingReport)
				err := daos.UpdateReport(report)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
					return
				}
				c.JSON(http.StatusCreated, map[string]string{"result": "success"})
			} else {
			}
		} else {
			fmt.Println("NotFound!!!")
		}
	}
	if hasMerchant == false {
		report.ID = bson.NewObjectId()
		report.IDMerchant = idMerchant
		report.Date = time.Now().Format("02-01-2006")
		report.ProductSelling = append(report.ProductSelling, productSellingReport)
		report.Accumulate = report.Accumulate + priceOfProductTotal

		if err := daos.InsertToReport(report); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusCreated, map[string]string{"result": "success"})
	}
}

func AllReportMerchantEndPoint(c *gin.Context) {

	reports, err := daos.FindAllReport()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helper.MapDataReport(reports))
}

func FindByIdReportMerchantEndPoint(c *gin.Context) {

	reports, err := daos.FindAllReport()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	var report model.Report
	var errR error
	for _, listReport := range reports {
		if listReport.IDMerchant == bson.ObjectIdHex(c.Param("id")) {
			report, errR = daos.FindByIdReport(listReport.ID)
			if errR != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
				return
			}
			c.JSON(http.StatusOK, helper.MapDataReport(report))
		} else {
			fmt.Println("NotFound!!!")
		}
	}
}
