package main

import (
	"kpay/middleware"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() *gin.Engine {

	r := gin.Default()
	merchant := r.Group("/merchant")
	merchant.Use(middleware.BasicAuthenMerchant)
	merchant.GET("/", AllMerchantEndPoint)
	merchant.POST("/register", CreateMerchantEndPoint)
	merchant.GET("/register/:id", FindByIdMerchantEndPoint)
	merchant.POST("/register/:id", UpdateIdMerchantEndPoint)

	merchant.GET("/register/:id/products", FindByIdMerchantEndPoint)
	merchant.POST("/register/:id/product", CreateProductMerchantEndPoint)
	merchant.DELETE("/register/:id/product/:product_id", DeleteProductMerchantEndPoint)
	merchant.POST("/register/:id/product/:product_id", UpdateProductMerchantEndPoint)
	merchant.GET("/register/:id/report", FindByIdReportMerchantEndPoint)
	merchant.GET("/report", AllReportMerchantEndPoint)

	buy := r.Group("/buy")
	buy.POST("/product", BuyProductInMerchantEndPoint)
	return r
}
