package mysql_api

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "chl"
	PASSWORD = "chl123"
	NETWORK  = "tcp"
	SERVER   = "192.168.184.130"
	PORT     = 3306
	DATABASE = "staff"
)

type Dbworker struct {
	Dsn			string
	Db			*sql.DB
	StaffInfo	StaffType
}

type StaffType struct {
	ID             int
	Name           sql.NullString
	Position       sql.NullString
	PhoneNumber    sql.NullString
}

func CreateTable(db *sql.DB, comm string) ret string, err error{
	stmt, _ := db.Prepare(comm)
	defer stmt.Close()

	if _, err := stmt.Exec(); err != nil {
		fmt.Printf("create table failed: %v\n", err)
		return _, err
	}
	fmt.Println("create table successed!!")
	return "create table successed", err
}

func InsertData(db *sql.DB, comm string) ret sql.Result, err error {
	stmt, _ := db.Prepare(comm)
	defer stmt.Close()

	if ret, err := stmt.Exec(); err != nil {
		fmt.Printf("insert data failed: %v\n", err)
		return _, err
	}
	fmt.Println("insert data successed!!")
	return ret, err
}

func QueryOne(db *sql.DB, string comm, param string) StaffType{}, error{
	stmt, _ := db.PrePare(comm)
	defer stmt.Close()

	row := stmt.QueryRow(param)
	var staffInfo StaffType{}
	if err  := row.Scan(&staffInfo.ID, &staffInfo.Name, &staffInfo.PhoneNumber, &staffInfo.Position); err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return staffInfo, err
	}

	return staffInfo, err
}

func QueryMulti(db *sql.DB, string comm) {

}

func UpdateData(db *sql.DB) {

}

func DeleteData(db *sql.DB) {

}