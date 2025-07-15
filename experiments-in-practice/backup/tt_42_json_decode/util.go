package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	// PepperSlave 主要資料庫
	PepperSlave *gorm.DB
)

// DBDetail DB的設定
type DBDetail struct {
	account  string // 帳號
	password string // 密碼
	ip       string // IP
	port     string // host
	dbname   string // 資料庫名稱
}

func init() {
	PepperSlave = setPepperSlave()

}

func setPepperSlave() *gorm.DB {
	d := DBDetail{
		account:  "root",
		password: "qwe123",
		ip:       "127.0.0.1",
		port:     "3306",
		dbname:   "fish",
	}
	con := d.concatDBConnectionStr()

	return createConnect(con)
}

func (d *DBDetail) concatDBConnectionStr() string {
	connectStr := d.account + ":" + d.password + "@tcp(" + d.ip + ":" + d.port
	connectStr += ")/" + d.dbname + "?parseTime=true"

	log.Println("連線字串是:", connectStr)

	return connectStr
}

func createConnect(con string) *gorm.DB {

	db, err := gorm.Open("mysql", con)
	if err != nil {
		log.Println(err)
	}

	// 限制連線數
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(30)
	db.DB().SetConnMaxLifetime(120 * time.Second)

	return db
}
