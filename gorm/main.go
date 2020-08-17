package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
	db, err := gorm.Open("mysql", con)
	if err != nil {
		log.Println("🐼🐼db open error:", err)
	}

	// 限制連線數
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(30)
	db.DB().SetConnMaxLifetime(120 * time.Second)

	db.SetLogger(&GormLogger{})


	return db, err
}

func main() {
	InitDB(EnvVariable{
		ExampleDBStr: "root:qwe123@tcp(127.0.0.1:3306)/example?charset=utf8&parseTime=True&loc=Local",
	})

	InsertRecord(6653, 10.5, "pika")


	GetTransRecord(6653,0,"")
}