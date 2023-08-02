package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	RecordMasterConn *gorm.DB
)

// EnvVariable 環境連線資訊
type EnvVariable struct {
	ExampleDBStr string
}

// InitDB 初始化DB連線
func InitDB(connStr EnvVariable) error {

	var dbConnErr error

	RecordMasterConn, dbConnErr = CreateGormInstance(connStr.ExampleDBStr)
	if dbConnErr != nil {
		log.Println("RecordMasterConn DB init  err:", dbConnErr)
		panic("db init error")
	}

	return dbConnErr
}

// CreateGormInstance 初始化 DB 連線實體
func CreateGormInstance(conn string) (*gorm.DB, error) {
	return createConnect(conn)
}

func createConnect(con string) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(con), &gorm.Config{})
	if err != nil {
		log.Println("🐼🐼db open error:", err)
	}
	dbInstance, dErr := db.DB()
	if err != nil {
		log.Println("🐼🐼db get instance error:", dErr)
	}
	// 限制連線數
	dbInstance.SetMaxIdleConns(5)
	dbInstance.SetMaxOpenConns(30)
	dbInstance.SetConnMaxLifetime(120 * time.Second)

	// customer logger 可能在gorm 2.0失效或者改method
	// 需要再找時間研究
	//dbInstance.SetLogger(&GormLogger{})

	return db, err
}
