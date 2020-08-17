package main

// GetTransRecord 取得紀錄
func GetTransRecord(userID, opcode int64, account string) []TableTransRecord {

	RecordMasterConn.LogMode(true)

	var record []TableTransRecord

	query := RecordMasterConn.Table("trans_record").Where(`user_id = ? `, userID,)

	if account != "" {
		query = query.Where("account = ?", account)
	}

	if opcode != 0 {
		query = query.Where("opcode = ?", opcode)
	}

	query.Find(&record)

	return record
}
