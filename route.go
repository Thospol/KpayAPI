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
	merchant.POST("/:id/report", CreateReportMerchantEndPoint)

	report := r.Group("/reports")
	report.GET("/", AllReportMerchantEndPoint)

	merchant.Use(middleware.BasicAuthenUser)
	buy := r.Group("/buy")
	buy.POST("/:id/product", BuyProductInMerchantEndPoint)

	user := r.Group("/user")
	user.POST("/", CreateUserEndPoint)
	user.GET("/", FindAllUserEndPoint)
	user.GET("/:id", FindByIDUserEndPoint)
	user.POST("/:id ", CreateBankAccountOfuserEndPoint)
	user.PUT("/:id", UpdateUserEndPoint)
	user.DELETE("/:id", DeleteUserEndPoint)
	user.DELETE("/:id/bankAccount", UpdateUserBankAccountEndPoint)
	return r
}
