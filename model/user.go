package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Identity      string    `json:"identity"`
	ClientIp      string    `json:"client_ip"`
	ClientPort    string    `json:"client_port"`
	LoginTime     time.Time `gorm:"autoUpdateTime:milli" json:"login_time"`
	HeartbeatTime time.Time `gorm:"autoUpdateTime:milli" json:"heartbeat_time"`
	LoginOutTime  time.Time `gorm:"autoUpdateTime:milli" json:"login_out_time""`
	IsLogOut      bool      `json:"is_log_out"`
	DeviceInfo    string    `json:"device_info"`
}

// 创建用户参数

type UserParams struct {
	Username string `gorm:"username;not null" json:"username"`
	Password string `gorm:"password;not null" json:"password"`
	Phone    string `gorm:"phone;not null" json:"phone"`
	Email    string `gorm:"email;not null" json:"email"`
}
