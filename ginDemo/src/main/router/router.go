package router

import (
	"ginDemo/src/main/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//r.GET("/test", dbUtils.AddTest)
	userGroup := r.Group("user")
	{
		userGroup.POST("Register", controller.Register)
		userGroup.POST("login", controller.Login)
	}
	return r
}
