package svc

import (
	"easy-chat/examples/user/api/internal/config"
	"easy-chat/examples/user/api/internal/middleware"
	"easy-chat/examples/user/rpc/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	userclient.User
	// 添加中件间
	LoginVerify rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		// 使用中间件
		LoginVerify: middleware.NewLoginVerifyMiddleware().Handle,
	}
}
