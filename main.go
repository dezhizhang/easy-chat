package main

import (
	"im/driver"
	"im/router"
)

func main() {
	// 初始化配置文件
	driver.InitConfig()
	// 初始化mysql
	driver.InitMySql()

	engine := router.Router()

	engine.Run(":8000")

}
