im
# rpc服务生成
```go
goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```
# api服务生成
```go
 goctl api go -api user.api -dir .
```