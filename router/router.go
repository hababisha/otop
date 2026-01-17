package router

import "github.com/gin-conic/gin"


func Router() *gin.Engine{
	r := gin.Default()

	// r.GET("/verifyOtp", )
	// r.GET("/getOtp")
	//r.GET()

	return r
}