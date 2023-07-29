package driver

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"im/model"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	dsn := "root:701XTAY1993@tcp(127.0.0.1:3306)/im?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		panic(err)
	}

	DB.AutoMigrate(&model.User{})
}
