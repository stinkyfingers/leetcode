package main

import (
	"log"
	"net/http"
	"time"

	"github.com/stinkyfingers/leetcode"
	"golang.org/x/net/websocket"
)

// TODO all tests

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/echo", websocket.Handler(handleEcho))
	mux.Handle("/time", websocket.Handler(handleTime))
	// TODO CORS

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleTime(ws *websocket.Conn) {
	// TODO handle exits gracefully
	ticker := time.Tick(time.Second * 10)
	for range ticker {
		now := time.Now()
		timeByte := []byte(now.String())
		_, err := ws.Write(timeByte)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func handleEcho(ws *websocket.Conn) {
	var m leetcode.Message
	err := websocket.JSON.Receive(ws, &m)
	if err != nil {
		e := Error{Code: 500, Message: err.Error()}
		websocket.JSON.Send(ws, e)
		return
	}

	err = websocket.JSON.Send(ws, m)
	if err != nil {
		e := Error{Code: 500, Message: err.Error()}
		websocket.JSON.Send(ws, e)
		return
	}
}
