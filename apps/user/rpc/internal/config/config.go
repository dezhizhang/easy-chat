package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	// mysql
	Mysql struct {
		DataSource string
	}
	// 缓存
	Cache cache.CacheConf

	// 鉴权
	Jwt struct {
		AccessSecret string
		AccessExpire int64
	}
}
