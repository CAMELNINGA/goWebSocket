package main

import (
	"goWebSocket/pkg/websocket"
	"log"
	"net/http"
)

import (
	"flag"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()
	hub := websocket.NewHub()
	go hub.Run()
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
