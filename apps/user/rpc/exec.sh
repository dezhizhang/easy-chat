goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.
goctl model mysql ddl -src="./user.sql" dir=. -c
