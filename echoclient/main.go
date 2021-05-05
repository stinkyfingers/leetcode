package main

import (
	"fmt"
	"io"
	"log"

	"github.com/stinkyfingers/leetcode"
	"golang.org/x/net/websocket"
)

func main() {
	origin := "http://localhost/"
	url := "ws://localhost:8080/echo"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	m := &leetcode.Message{
		Data: "this is a test",
	}

	err = websocket.JSON.Send(ws, m)
	if err != nil {
		log.Fatal(err)
	}

	for {
		var msg = make([]byte, 512)
		var n int
		n, err = ws.Read(msg)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}
		fmt.Printf("Received: %s.\n", msg[:n])
	}
}
