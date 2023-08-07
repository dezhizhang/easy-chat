package service

import (
	"im/driver"
	"im/model"
)

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
