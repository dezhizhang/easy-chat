package websocket

import "github.com/gorilla/websocket"

type Route struct {
	Method  string      `json:"method"`
	Handler HandlerFunc `json:"handler"`
}

type HandlerFunc func(srv *Server, conn *websocket.Conn, msg *Message)
