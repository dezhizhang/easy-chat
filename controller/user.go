package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取user成功"})
}
