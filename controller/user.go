package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取user成功"})
}
