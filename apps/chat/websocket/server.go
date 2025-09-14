package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"sync"
)

type Server struct {
	add string
	logx.Logger
	sync.RWMutex
	authentication Authentication
	upgrader       websocket.Upgrader
	routes         map[string]HandlerFunc
	connToUser     map[*websocket.Conn]string
	userToConn     map[string]*websocket.Conn
}

func NewServer(add string) *Server {
	return &Server{
		add:            add,
		upgrader:       websocket.Upgrader{},
		authentication: new(authentication),
		routes:         make(map[string]HandlerFunc),
		connToUser:     make(map[*websocket.Conn]string),
		userToConn:     make(map[string]*websocket.Conn),
		Logger:         logx.WithContext(context.Background()),
	}
}

func (s *Server) ServerWs(w http.ResponseWriter, r *http.Request) {
	// 处理服务异常
	defer func() {
		if err := recover(); err != nil {
			s.Error("server handler panic:", err)
		}
	}()

	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.Errorf("upgrader error: %v", err)
		return
	}

	// 根据连接对像获取请求信息
	go s.HandlerConn(conn)
}

// AddConn 添加连接池
func (s *Server) AddConn(conn *websocket.Conn, req *http.Request) {
	uid := s.authentication.UserId(req)
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()
	s.connToUser[conn] = uid
	s.userToConn[uid] = conn
}

// GetConn 获取连接对像
func (s *Server) GetConn(uid string) *websocket.Conn {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()
	return s.userToConn[uid]
}

// HandlerConn 根据连接对像执行任务处理
func (s *Server) HandlerConn(conn *websocket.Conn) {
	for {

		_, msg, err := conn.ReadMessage()
		if err != nil {
			s.Errorf("websocket conn message error: %v", err)
			//todo
			break
		}

		var message Message
		if err = json.Unmarshal(msg, &message); err != nil {
			s.Errorf("websocket conn message error: %v", err)
			break
		}
		// 根据请求的method分发路由并执行
		if handler, ok := s.routes[message.Method]; ok {
			handler(s, conn, &message)
		} else {
			conn.WriteMessage(websocket.CloseMessage, []byte(fmt.Sprintf("不存在执行的方法%v ", message.Method)))
		}

	}
}

// AddRoutes 添加路由方法
func (s *Server) AddRoutes(rs []Route) {
	for _, r := range rs {
		s.routes[r.Method] = r.Handler
	}
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
