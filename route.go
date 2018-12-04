package main

import "github.com/gin-gonic/gin"

func initializeRoutes() *gin.Engine {

	r := gin.Default()
	merchant := r.Group("/merchant")
	merchant.GET("/", AllMerchantEndPoint)
	merchant.POST("/register", CreateMerchantEndPoint)
	merchant.GET("/register/:id", FindByIdMerchantEndPoint)
	merchant.POST("/register/:id", UpdateIdMerchantEndPoint)

	merchant.GET("/register/:id/products", FindByIdMerchantEndPoint)
	merchant.POST("/register/:id/product", CreateProductMerchantEndPoint)
	merchant.DELETE("/register/:id/product/:product_id", DeleteProductMerchantEndPoint)
	return r
}
