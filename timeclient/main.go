package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

func main() {
	origin := "http://localhost/"
	url := "ws://localhost:8080/time"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	for {
		var msg = make([]byte, 512)
		var n int
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received: %s.\n", msg[:n])
	}
}
