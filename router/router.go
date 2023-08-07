package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"im/controller"
	"im/docs"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	v1 := r.Group("/api/v1")
	{
		v1.GET("/user", controller.GetUser)
		v1.POST("/user/create", controller.CreateUser)
		v1.GET("/user/delete", controller.DeleteUser)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
