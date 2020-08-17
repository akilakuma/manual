package main

import "time"

// TableTransRecord 交易紀錄資料表格式
type TableTransRecord struct {
	ID        int64   `gorm:"column:id"`         // 流水號
	UserID    int64   `gorm:"column:user_id"`    // user_id
	Account   string  `gorm:"column:account"`    // 使用者名稱
	Point     float64 `gorm:"column:point"`      // 交易金額
	OPcode    int64   `gorm:"column:opcode"`     // opcode
	CreatedAt int64   `gorm:"column:created_at"` // 建立時間
}

// TableName TableTransRecord 在DB的資料表名稱
func (t *TableTransRecord) TableName() string {
	return "trans_record"
}

// InsertRecord 寫入紀錄
func InsertRecord(userID int64, amount float64, account string) {

	RecordMasterConn.LogMode(true)

	var record = TableTransRecord{
		UserID:  userID,
		Account: "bananaKing",
		Point:   amount,
		OPcode: 1001,
		CreatedAt: time.Now().Unix(),
	}
	RecordMasterConn.Create(&record)

	// TODO error report
}
