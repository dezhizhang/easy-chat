package config

import "github.com/zeromicro/go-zero/core/service"

type Config struct {
	service.ServiceConf
	// 听地址
	ListenOn string
}
