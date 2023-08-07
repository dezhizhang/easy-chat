package service

import (
	"fmt"
	"im/driver"
	"im/model"
)

func GetUser() (*[]model.User, error) {
	var user []model.User
	err := driver.DB.Model(&user).Find(&user).Error
	return &user, err

}

func CreateUser(params *model.UserParams) error {
	user := model.User{
		Email:    params.Email,
		Phone:    params.Phone,
		Username: params.Username,
		Password: params.Password,
	}
	err := driver.DB.Model(&model.User{}).Create(&user).Error
	return err
}

func DeleteUser(userid int) error {
	//db.Where("name = ?", "jinzhu").Delete(&email)
	var user model.User
	err := driver.DB.Model(&user).Where("id = ?", userid).Delete(&user).Error
	fmt.Println("err", err)
	return err
}
