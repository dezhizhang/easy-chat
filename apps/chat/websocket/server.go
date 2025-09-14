package websocket

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type Server struct {
	add      string
	upgrader websocket.Upgrader
	logx.Logger
}

func NewServer(add string) *Server {
	return &Server{
		add:      add,
		upgrader: websocket.Upgrader{},
		Logger:   logx.WithContext(context.Background()),
	}
}

func (s *Server) ServerWs(w http.ResponseWriter, r *http.Request) {

}

// Start 启动服务
func (s *Server) Start() {
	http.HandleFunc("/ws", s.ServerWs)
	fmt.Printf("websocket server listening at %s\n", s.add)
	s.Info(http.ListenAndServe(s.add, nil))
}

func (s *Server) Stop() {
	fmt.Println("停止服务")
}
