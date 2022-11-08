//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

var addr = flag.String("addr", "192.168.56.103:30080", "http service address")
//var addr = flag.String("addr", "127.0.0.1:8080", "http service address")

func main() {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws", RawQuery: "versionID=ec-builder&tail=-1&follow=true"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			fmt.Println(string(message))
		}
	}()
	<-done
}
