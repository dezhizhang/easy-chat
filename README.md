# easy-chat

## rpc服务与api服务的创建

### 1. rpc的protobuf文件创建

```protobuf
syntax = "proto3";

package user;

option go_package = "./user";


message  UserReq {
  string id = 1;
}

message UserResp {
  string id = 1;
  string name = 2;
  string phone = 3;
}

service User{
  rpc GetUser(UserReq) returns(UserResp);
}
```

### 2. rpc服务的创建

```bash
goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.  #通过proto生成服务
```

### 3. api的protobuf文件创建

### api服务的创建

```bash
goctl api new api # 生成
goctl api go -api user.api -dir . style=gozero

```

## api 服务调用rpc服务

### 1. api服务etc的添加

```yaml
Name: User
Host: 0.0.0.0
Port: 8888

# userRpc服务调用
UserRpc:
  Etcd:
    Hosts:
      - 127.0.0.1:2379
    Key: user.rpc

```

### 2. api服务配置的添加

```go
package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf // 添加调用rpc服务
}

```

### 3. api服务svc配置请求上下文

```go
package svc

import (
	"easy-chat/examples/user/api/internal/config"
	"easy-chat/examples/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}

```

### 4. api服务逻辑层调用rpc服务
```go
// GetUser 获取用户信息
func (l *GetUserLogic) GetUser(req *types.UserReq) (resp *types.UserResp, err error) {

	user, err := l.svcCtx.User.GetUser(l.ctx, &userclient.UserReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.UserResp{
		Id:    user.Id,
		Name:  user.Name,
		Phone: user.Phone,
	}, nil
}

```

## rpc 服务响应api服务

### 1. 逻辑层响应api服务
```go
// GetUser rpc 服务处理逻辑
func (l *GetUserLogic) GetUser(in *user.UserReq) (*user.UserResp, error) {

return &user.UserResp{
Id:    "123456",
Name:  "tom",
Phone: "159924****8",
}, nil
}

```
