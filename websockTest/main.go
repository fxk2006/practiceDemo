package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
)
var upgrader = websocket.Upgrader{}


func main() {
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(":9090", nil))
}


func echo( w http.ResponseWriter, r *http.Request){
	c, _ := upgrader.Upgrade(w, r, nil)
	defer c.Close()
	for {
		mt, _, _ := c.ReadMessage()
		fmt.Println("收到：")
		c.WriteMessage(mt,[]byte("hello"))
	}
}