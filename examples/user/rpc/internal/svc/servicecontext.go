package svc

import (
	"easy-chat/examples/user/model"
	"easy-chat/examples/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// mysql 连接信息
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		// mysql和redis配置
		UserModel: model.NewUsersModel(conn, c.Cache),
	}
}
