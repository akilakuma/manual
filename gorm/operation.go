package main

import (
	"gorm.io/gorm/clause"
	"time"
)

// TableTransRecord 交易紀錄資料表格式
// 如果insert之後，想取得auto increment 產生的id，需要上gorm tag (如果命名`id`好像會自動識別的樣子)
type TableTransRecord struct {
	ID        int64   `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY;not null"` // 流水號
	UserID    int64   `gorm:"column:user_id"`                                // user_id
	Account   string  `gorm:"column:account"`                                // 使用者名稱
	Point     float64 `gorm:"column:point"`                                  // 交易金額
	OPcode    int64   `gorm:"column:opcode"`                                 // opcode
	CreatedAt int64   `gorm:"column:created_at"`                             // 建立時間
}

// TableName TableTransRecord 在DB的資料表名稱
func (t *TableTransRecord) TableName() string {
	return "trans_record"
}

/* ========================== insert 範例 ==========================*/

// InsertRecord 寫入紀錄
func InsertRecord(userID int64, amount float64, account string) {

	RecordMasterConn.Debug()

	var record = TableTransRecord{
		UserID:    userID,
		Account:   "bananaKing",
		Point:     amount,
		OPcode:    1001,
		CreatedAt: time.Now().Unix(),
	}
	RecordMasterConn.Create(&record)

	// TODO error report
}

// BulkInsertDuplicateUpdate 寫入紀錄
// 如果insert遇到duplicate 可以透過指定的方式，update 想要的欄位
func BulkInsertDuplicateUpdate(data []TableTransRecord) {

	// 執行 Bulk Insert
	RecordMasterConn.Table("trans_record").
		Clauses(clause.OnConflict{
			//Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"point"}),
		}).Create(&data)

}

/*

// 根据 map 更新
db.Table("users").Where("id IN ?", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);
*/

/* ========================== select 範例 ==========================*/

// GetTransRecord 取得紀錄
func GetTransRecord(userID, opcode int64, account string) []TableTransRecord {

	RecordMasterConn.Debug()

	var record []TableTransRecord

	query := RecordMasterConn.Table("trans_record").Where(`user_id = ? `, userID)

	if account != "" {
		query = query.Where("account = ?", account)
	}

	if opcode != 0 {
		query = query.Where("opcode = ?", opcode)
	}

	query.Find(&record)

	return record
}
