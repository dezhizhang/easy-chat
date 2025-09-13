package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	// mysql配置
	Mysql struct {
		DataSource string
	}
	// redis配置
	Cache cache.CacheConf
}
