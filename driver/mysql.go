package driver

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"im/model"
)

var (
	DB *gorm.DB
)

func InitMySql() {
	var err error
	port := viper.GetInt("mysql.port")
	user := viper.GetString("mysql.user")
	host := viper.GetString("mysql.host")
	dbname := viper.GetString("mysql.dbname")
	password := viper.GetString("mysql.password")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)
	fmt.Println("dsn", dsn)

	//dsn := "root:701XTAY1993@tcp(127.0.0.1:3306)/im?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&model.User{})
}
