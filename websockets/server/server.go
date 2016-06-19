package main

import (
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Listening...")
	http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
	os.Exit(0)
}

func EchoServer(ws *websocket.Conn) {
	log.Print("Got echo request")
	io.Copy(ws, ws)
}
