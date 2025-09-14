package test

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"testing"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域
	},
}

func ServerWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}

	}

}

func TestWebsocket(t *testing.T) {
	http.HandleFunc("/ws", ServerWs)
	fmt.Println("启动 websocket 服务，地址: ws://127.0.0.1:9090/ws")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
