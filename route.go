package main

import (
	"kpay/middleware"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() *gin.Engine {

	r := gin.Default()
	register := r.Group("/register")
	register.POST("/", CreateMerchantEndPoint)
	merchants := r.Group("/merchants")
	merchants.GET("/", AllMerchantEndPoint)

	merchant := r.Group("/merchant")
	merchant.Use(middleware.BasicAuthenMerchant)
	merchant.GET("/:id", FindByIdMerchantEndPoint)
	merchant.POST("/:id", UpdateIdMerchantEndPoint)

	merchant.GET("/:id/products", FindByIdProductMerchantEndPoint)
	merchant.POST("/:id/product", CreateProductMerchantEndPoint)
	merchant.DELETE("/:id/product/:product_id", DeleteProductMerchantEndPoint)
	merchant.POST("/:id/product/:product_id", UpdateProductMerchantEndPoint)
	merchant.GET("/:id/report", FindByIdReportMerchantEndPoint)

	report := r.Group("/report")
	report.GET("/", AllReportMerchantEndPoint)

	buy := r.Group("/buy")
	buy.POST("/product", BuyProductInMerchantEndPoint)
	return r
}
