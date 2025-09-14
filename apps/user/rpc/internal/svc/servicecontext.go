package svc

import (
	"easy-chat/apps/user/model"
	"easy-chat/apps/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// 用户模型
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		// 创建mysql连接
		UserModel: model.NewUsersModel(conn, c.Cache),
	}
}
