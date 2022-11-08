//go:build ignore
// +build ignore
//
//package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/hpcloud/tail"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	// Time allowed to write the file to the client.
	//writeWait = 1 * time.Second
	writeWait = 100 * time.Millisecond

	// Time allowed to read the next pong message from the client.
	//pongWait = 24 * time.Hour
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

var (
	filename = "project/network/websocket/logfile"
	addr     = flag.String("addr", ":8080", "http service address")
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func reader(ws *websocket.Conn) {
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

func tailFile() *tail.Tail {
	tailfs, err := tail.TailFile(filename, tail.Config{
		ReOpen:    false,                                // 文件被移除或被打包，需要重新打开
		Follow:    true,                                 // 实时跟踪
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 如果程序出现异常，保存上次读取的位置，避免重新读取。
		MustExist: false,                                // 如果文件不存在，是否推出程序，false是不退出
		Poll:      true,
	})

	if err != nil {
		fmt.Println("tailf failed, err:", err)
		return nil
	}
	return tailfs
}

func writer(ws *websocket.Conn) {
	tailfs := tailFile()
	pingTicker := time.NewTicker(pingPeriod)
	//ticker := time.NewTicker(20*time.Second)
	defer func() {
		pingTicker.Stop()
		ws.Close()
		log.Println("connection closed")
	}()

	for {
		select {
		case msg, ok := <-tailfs.Lines:
			if ok {
				ws.SetWriteDeadline(time.Now().Add(writeWait))
				fmt.Printf("read file content： %s\n", msg)
				if err := ws.WriteMessage(websocket.TextMessage, []byte(msg.Text)); err != nil {
					return
				}
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	q := r.URL.Query()
	versionID := q.Get("versionID")
	fmt.Printf(versionID)

	ori, err := ioutil.ReadFile(filename)
	if err != nil {
		ori = []byte(err.Error())
	}

	err = ws.WriteMessage(websocket.TextMessage, ori)
	if err != nil {
		log.Println("write:", err)
		return
	}

	go writer(ws)
	//reader(ws)
}

func main() {
	http.HandleFunc("/ws", serveWs)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
