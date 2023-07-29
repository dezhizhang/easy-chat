package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string `json:"name"`
	Password      string `json:"password"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Identity      string `json:"identity"`
	ClientIp      string `json:"clientIp"`
	ClientPort    string `json:"clientPort"`
	LoginTime     uint64 `json:"loginTime"`
	HeartbeatTime uint64 `json:"heartbeatTime"`
	LogOutTime    uint64 `json:"logOutTime"`
	IsLogOut      bool   `json:"isLogOut"`
	DeviceInfo    string `json:"deviceInfo"`
}
