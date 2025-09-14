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

## 数据库的操作

### 1. 编写操作的sql文件并生成对应的操作数据库接口

```bash
CREATE TABLE `users`
(
    `id`         varchar(24) COLLATE utf8mb4_unicode_ci  NOT NULL,
    `avatar`     varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `name`       varchar(24) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '',
    `phone`      varchar(20) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '',
    `password`   varchar(191) COLLATE utf8mb4_unicode_ci          DEFAULT NULL,
    `status`     int(10) DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

```bash
# -c是增加缓存
goctl model mysql ddl -src user.sql -dir . -c 
```

### 2. proto文件增加添加用户接口

```protobuf
syntax = "proto3";

package user;

option go_package = "./user";


message CreateReq {
  string id = 1;
  string name = 2;
  string phone = 3;
}

message CreateResp {
  string msg = 1;
}

message  UserReq {
  string id = 1;
}

message UserResp {
  string id = 1;
  string name = 2;
  string phone = 3;
}

service User{
  // CreateUser 创建用户信息
  rpc CreateUser(CreateReq) returns(CreateResp);
  //GetUser  获取用户信息
  rpc GetUser(UserReq) returns(UserResp);

}
```

### 3. 生成对应的rpc接口

```bash
goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.  #通过proto生成服务
```

### 4. 添加etc请求配置

```yaml
Name: user.rpc
ListenOn: 0.0.0.0:8080

Etcd:
  Hosts:
    - 127.0.0.1:2379
  Key: user.rpc

Mysql:
  DataSource: root:V4Nn9fa#Xf!@tpc(127.0.0.1:3306)/user?charset=utf8mb4

Cache:
  - Host: 127.0.0.1:6379
    Type: node
    Pass:

```

### 5. 添加配置结构体

```go
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

```

### 6. 添加svc 上下文

```go
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
```

### 7. logic增加用户逻辑

```go
// CreateUser  创建用户信息
func (l *CreateUserLogic) CreateUser(in *user.CreateReq) (*user.CreateResp, error) {
_, err := l.svcCtx.UserModel.Insert(l.ctx, &model.Users{
Id:    in.Id,
Name:  in.Name,
Phone: in.Phone,
})

if err != nil {
return nil, err
}

return &user.CreateResp{Msg: "ok"}, nil
}

```

## api 服务调用rpc服务

### 1. 增加添加用户api

```api
syntax = "v1"

type CreateReq {
    Id string `json:"id"`
    Name string `json:"name"`
    Phone string `json:"phone"`
}

type CreateResp {
    Msg string `json:"msg"`
}

type UserReq {
    Id string `json:"id"`
}

type UserResp {
    Id string `json:"id"`
    Name string `json:"name"`
    Phone string `json:"phone"`
}

service User {
    // createUser 添加用户
    @handler createUser
    post /user (CreateReq) returns (CreateResp)

    // getUser 获取用户
    @handler getUser
    get /user (UserReq) returns (UserResp)
}

```

### 2. 生成对应的api服务接口

```bash
goctl api go -api user.api -dir . style = gozero
```

### 3. 编写对应的逻辑
```go

func (l *CreateUserLogic) CreateUser(req *types.CreateReq) (resp *types.CreateResp, err error) {
	user, err := l.svcCtx.User.CreateUser(l.ctx, &userclient.CreateReq{
		Id:    req.Id,
		Name:  req.Name,
		Phone: req.Phone,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResp{Msg: user.Msg}, nil
}
```
## mongodb的生成
```bash
goctl model mongo -type User -c dir .
```