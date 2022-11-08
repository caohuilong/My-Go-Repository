package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
const (
	USERNAME = "chl"
	PASSWORD = "chl123"
	NETWORK  = "tcp"
	SERVER   = "192.168.184.130"
	PORT     = 3306
	DATABASE = "staff"
)*/

type dbWorker struct {
	Dsn       string
	Db        *sql.DB
	StaffInfo staffTB
}

type staffTB struct {
	ID             int
	Name           sql.NullString
	Position       sql.NullString
	PhoneNumber    sql.NullString
	OnboardingDate sql.NullTime
}

func main() {
	var err error
	dbw := dbWorker{
		//Dsn: USERNAME + ":" + PASSWORD + "@" + NETWORK + "(" + SERVER + ":" + PORT + ")/" + DATABASE + "?charset=utf8mb4",
		Dsn: "chl:chl123@tcp(192.168.184.130:3306)/staff?charset=utf8mb4",
	}
	dbw.Db, err = sql.Open("mysql", dbw.Dsn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}
	defer dbw.Db.Close()

	dbw.createTable()
	dbw.insertData()
	dbw.queryOne()
	dbw.queryMulti()

	dbw.updateData()
	dbw.queryMulti()

	dbw.deleteData()
	dbw.queryMulti()
}

func (dbw *dbWorker) createTable() {
	stmt, _ := dbw.Db.Prepare(`CREATE TABLE IF NOT EXISTS uestc_staff(
        id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(40) NOT NULL,
        position TEXT NOT NULL,
        phone_number VARCHAR(11) NOT NULL,
        onboarding_date DATE)`)
	defer stmt.Close()

	if _, err := stmt.Exec(); err != nil {
		fmt.Println("create table failed: ", err)
		return
	}
	fmt.Println("create table successed")
}

func (dbw *dbWorker) insertData() {
	stmt, _ := dbw.Db.Prepare(`INSERT INTO uestc_staff (name, position, phone_number, onboarding_date) VALUES (?, ?, ?, ?)`)
	defer stmt.Close()

	ret, err := stmt.Exec("wangdan", "后端开发工程师", "17328681265", "2020.07.07")
	if err != nil {
		fmt.Printf("insert data err:%v\n", err)
		return
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:", RowsAffected)
	}

	ret, err = stmt.Exec("caohuilng", "后端开发工程师", "17328680796", "2020.07.07")
	if err != nil {
		fmt.Printf("insert data err:%v\n", err)
		return
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:", RowsAffected)
	}
}

//查询单行
func (dbw *dbWorker) queryOne() {
	stmt, _ := dbw.Db.Prepare("select * from uestc_staff where name=?")
	stmt.Close()

	row := stmt.QueryRow("wangdan")
	dbw.StaffInfo = staffTB{}
	if err := row.Scan(&dbw.StaffInfo.ID, &dbw.StaffInfo.Name, &dbw.StaffInfo.Position, &dbw.StaffInfo.PhoneNumber, &dbw.StaffInfo.OnboardingDate); err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Println("get data, id: ", dbw.StaffInfo.ID, "\nname: ", dbw.StaffInfo.Name, "\nposition: ", dbw.StaffInfo.Position, "\nphone_number: ", dbw.StaffInfo.PhoneNumber, "\nonboarding_date: ", dbw.StaffInfo.OnboardingDate)
}

//查询多行
func (dbw *dbWorker) queryMulti() {
	stmt, _ := dbw.Db.Prepare(`SELECT * From uestc_staff`)
	defer stmt.Close()

	dbw.StaffInfo = staffTB{}

	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	for rows.Next() {
		rows.Scan(&dbw.StaffInfo.ID, &dbw.StaffInfo.Name, &dbw.StaffInfo.OnboardingDate, &dbw.StaffInfo.PhoneNumber, &dbw.StaffInfo.Position)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		if !dbw.StaffInfo.OnboardingDate.Valid {
			dbw.StaffInfo.OnboardingDate.Time = time.Now()
		}
		fmt.Println("get data, id: ", dbw.StaffInfo.ID, "\nname: ", dbw.StaffInfo.Name, "\nposition: ", dbw.StaffInfo.Position, "\nphone_number: ", dbw.StaffInfo.PhoneNumber, "\nonboarding_date: ", dbw.StaffInfo.OnboardingDate)
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func (dbw *dbWorker) updateData() {
	stmt, _ := dbw.Db.Prepare(`UPDATE uestc_staff set phone_number=? where name=?`)
	defer stmt.Close()

	ret, err := stmt.Exec("17328680796", "caohuilong")
	if err != nil {
		fmt.Printf("Update failed, err: %v\n", err)
		return
	}
	fmt.Println("update data successed: ", ret)

	RowsAffected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Println("Affected rows: ", RowsAffected)
}

func (dbw *dbWorker) deleteData() {
	stmt, _ := dbw.Db.Prepare(`DELETE FROM uestc_staff where name=?`)
	defer stmt.Close()

	ret, err := stmt.Exec("wangdan")
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}
	fmt.Println("delete data successed:", ret)

	RowsAffected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Println("Affected rows: ", RowsAffected)
}
