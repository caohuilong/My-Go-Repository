package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strconv"
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

func writer(ws *websocket.Conn, args []string) {
	defer func() {
		ws.Close()
		log.Println("connection closed")
	}()
	cmd := exec.Command("tail", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("tail failed, err: ", err)
		return
	}

	ws.SetWriteDeadline(time.Now().Add(writeWait))
	err = ws.WriteMessage(websocket.TextMessage, out)
	if err != nil {
		fmt.Println("write message failed, err: ", err)
		return
	}
}

func streamWriter(ws *websocket.Conn, args []string) {
	cmd := exec.Command("tail", args...)
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("tail failed, err: ", err)
		return
	}
	pingTicker := time.NewTicker(pingPeriod)
	ticker := time.NewTicker(20*time.Second)
	defer func() {
		pingTicker.Stop()
		ws.Close()
		log.Println("connection closed")
	}()

	reader := bufio.NewReader(stdoutPipe)

	err = cmd.Start()
	if err != nil {
		log.Println("cmd start failed, err: ", err)
		return
	}

	for {
		select {
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		case <-ticker.C:
			log.Println("task is over, close the websocket")
			return
		default:
			for reader.Size() > 0 {
				bytes, err := reader.ReadBytes('\n')
				if err != nil || err == io.EOF {
					log.Println("read bytes failed, err: ", err)
					return
				}
				ws.SetWriteDeadline(time.Now().Add(writeWait))
				err = ws.WriteMessage(websocket.TextMessage, bytes[:len(bytes)-1])
				if err != nil {
					log.Println("write message failed, err: ", err)
					return
				}
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
	defer func() {
		if err != nil {
			ws.Close()
		}
	}()

	log.Printf("connection, %v", r)

	q := r.URL.Query()
	//versionID := q.Get("versionID")
	tailStr := q.Get("tail")
	followStr := q.Get("follow")

	args := []string{}

	var tailNum int
	if tailStr == "" {
		tailNum = -1
	} else {
		tailNum, err = strconv.Atoi(tailStr)
		if err != nil {
			logrus.Error(err)
			return
		}
		if tailNum < -1 {
			err = errors.New("invalid tail")
			logrus.Error("invalid tail")
			return
		}
	}
	args = append(args, "-n")
	if tailNum == -1 {
		args = append(args, "+0")
	} else {
		args = append(args, strconv.Itoa(tailNum))
	}

	follow := false
	if followStr == "true" || followStr == "True" {
		follow = true
		args = append(args, "-f")
	}

	args = append(args, filename)
	fmt.Printf("%v", args)

	if !follow {
		go writer(ws, args)
	} else {
		go streamWriter(ws, args)
	}
	//go writer(ws, args)
	//reader(ws)
}

func main() {
	http.HandleFunc("/ws", serveWs)
	log.Println("listen and serve")
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
