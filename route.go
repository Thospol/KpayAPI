package main

import "github.com/gin-gonic/gin"

func initializeRoutes() *gin.Engine {

	r := gin.Default()
	merchant := r.Group("/merchant")
	merchant.POST("/register", CreateMerchantEndPoint)

	return r
}
