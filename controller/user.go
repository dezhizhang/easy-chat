package controller

import (
	"github.com/gin-gonic/gin"
	"im/model"
	"im/service"
	"log"
	"net/http"
	"strconv"
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
	user, err := service.GetUser()
	if err != nil {
		log.Fatalf("获取用户失败:%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取user成功", "data": user})
}

func CreateUser(c *gin.Context) {
	var userParams model.UserParams
	err := c.ShouldBindJSON(&userParams)
	if err != nil {
		log.Fatalf("获取参数失败:%s", err.Error())
		return
	}

	err = service.CreateUser(&userParams)
	if err != nil {
		log.Fatalf("创建用户失败:%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "创建用户成功", "data": nil})
}

func DeleteUser(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Query("id"))
	err := service.DeleteUser(userid)
	if err != nil {
		log.Fatalf("删除用户失败:%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "删除用户成功"})
}
