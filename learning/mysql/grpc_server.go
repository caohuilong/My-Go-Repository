package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user = "chl"
	password = "chl123"
	protocol = "tcp"
	addr = "192.168.184.130"
	port = "3306"
	database = "staff"
)

type Server struct {
	dbAddr		string
	driver		string
	dbconn		*sql.DB
}

func (srv *Server) Conn() error {
	db, err := sql.Open(srv.driver, srv.dbAddr)
	if err != nil {
		fmt.Printf("connection to mysql failed: %v\n", err)
		return err
	}
	srv.dbconn = db
	return srv.dbconn.Ping()
}

func (srv *Server) 

func main() {
	var server Server
	server.dbAddr = fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8", user, password, protocol, addr, port, database)
	server.driver = "mysql"
	if err := server.Conn(); err != nil {
		panic(err)
		return
	}
	
}