package main

import (
	"easy-chat/apps/chat/internal/config"
	"easy-chat/apps/chat/internal/svc"
	"easy-chat/apps/chat/websocket"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}

	svc.NewServiceContext(c)

	srv := websocket.NewServer(c.ListenOn)

	fmt.Printf("start websocket on %s\n", c.ListenOn)

	srv.Start()

}

//func main() {
//	flag.Parse()
//
//	var c config.Config
//	conf.MustLoad(*configFile, &c)
//
//	server := rest.MustNewServer(c.RestConf)
//	defer server.Stop()
//
//	ctx := svc.NewServiceContext(c)
//	handler.RegisterHandlers(server, ctx)
//
//	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
//	server.Start()
//}
