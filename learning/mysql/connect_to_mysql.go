package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type DbWorker struct {
    Dsn      string
    Db       *sql.DB
    UserInfo userTB
}

type UserTB struct {
    Id     int
    Name   sql.NullString
}

func main() {
    var err error
    dbw := DbWorker{
        Dsn: "chl:chl123@tcp(localhost:3306)/sqlx_db?charset=utf8mb4",
    }
    dbw.Db, err = sql.Open("mysql", dbw.Dsn)
    if err != nil {
        panic(err)
        return
    }
    defer dbw.Db.Close()

    dbw.insertData()
    dbw.queryData()
}

func (dbw *DbWorker) insertData() {
}
