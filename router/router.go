package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hababisha/otop/controller"
)

func Router() *gin.Engine{
	r := gin.Default()


	otp := r.Group("/otp")
	{
		otp.POST("/generate", controller.GenerateOTP)
		// otp.POST("/verify", controller.VerifyOTP)
	}
	return r
}