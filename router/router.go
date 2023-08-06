package router

import (
	"github.com/gin-gonic/gin"
	"im/controller"
)

func Router() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/user", controller.GetUser)
	}
	return r
}
