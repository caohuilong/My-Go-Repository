//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	filename = "project/network/websocket/logfile"
	addr     = flag.String("addr", ":8080", "http service address")
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}
	defer ws.Close()

	versionID := r.Header["Versionid"]
	fmt.Println(versionID)

	ori, err := ioutil.ReadFile(filename)
	if err != nil {
		ori = []byte(err.Error())
	}

	err = ws.WriteMessage(websocket.TextMessage, ori)
	if err != nil {
		log.Println("write:", err)
		return
	}

	filewatch(ws)
}

func filewatch(ws *websocket.Conn) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					ori, err := ioutil.ReadFile(filename)
					if err != nil {
						ori = []byte(err.Error())
					}

					err = ws.WriteMessage(websocket.TextMessage, ori)
					if err != nil {
						log.Println("write:", err)
						return
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(filename)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func main() {
	http.HandleFunc("/ws", serveWs)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
