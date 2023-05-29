package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	OrderDBStr *gorm.DB
)

// EnvVariable 環境連線資訊
type EnvVariable struct {
	OrderDBStr string
}

// InitDB 初始化DB連線
func InitDB(connStr EnvVariable) error {

	var dbConnErr error

	OrderDBStr, dbConnErr = CreateGormInstance(connStr.OrderDBStr)
	if dbConnErr != nil {
		log.Println("OrderDB DB init  err:", dbConnErr)
		panic("db init error")
	}

	return dbConnErr
}

// CreateGormInstance 初始化 DB 連線實體
func CreateGormInstance(conn string) (*gorm.DB, error) {
	return createConnect(conn)
}

func createConnect(con string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", con)
	if err != nil {
		log.Println("🐼🐼db open error:", err)
	}

	// 限制連線數
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(50)
	db.DB().SetConnMaxLifetime(1200 * time.Second)

	return db, err
}
