package driver

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败:%s", err.Error())
		return
	}
	fmt.Println("读取配置文件成功", viper.GetString("app_name"))
}
