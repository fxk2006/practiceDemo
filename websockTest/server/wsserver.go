package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

type WsServer struct {
	listener net.Listener
	addr     string
	upgrade  *websocket.Upgrader
}

func NewWsServer() *WsServer {
	ws := new(WsServer)
	ws.addr = "0.0.0.0:10215"
	ws.upgrade = &websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return ws
}

func (self *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	conn, err := self.upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("websocket error:", err)
		return
	}
	for {
		mty, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("收到消息：", string(msg))
		conn.WriteMessage(mty, []byte("hello websocket test"))
	}
}

func (w *WsServer) Start() (err error) {
	w.listener, err = net.Listen("tcp", w.addr)
	if err != nil {
		fmt.Println("net listen error:", err)
		return
	}

	err = http.Serve(w.listener, w)
	if err != nil {
		fmt.Println("http serve error:", err)
		return
	}

	return nil
}
