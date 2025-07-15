package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 一個在DB叫做User的table結構
type User struct {
	ID        uint   `gorm:"primary_key"`
	PartnerID int    `gorm:"column:partner_id";`
	Name      string `gorm:"column:name;"`
}

func main() {
	test1()
}

func test1() {
	// 建立連線
	// [root] 帳號，
	// [qwe123] 密碼
	// [tcp(127.0.0.1:3306)] ip和port而且一定要外刮一層tcp
	// [Match] DB的名稱
	db, err := gorm.Open("mysql", "root:qwe123@tcp(127.0.0.1:3306)/Test?charset=utf8&parseTime=True&loc=Local")
	// 選擇要不要打開debug mode 超實用
	// LogMode set log mode, `true` for detailed logs, `false` for no log, default, will only print error logs
	db.LogMode(true)

	// 連不連得到，這個超級重要!!
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	// Orm，跟db對起來，gorm用DB必先宣告
	db.AutoMigrate(&User{})

	// 試寫入一筆
	// u1 := User{PartnerID: 2, Name: "abc@hotmail.com"}
	// db.Create(&u1)

	// 讀取剛剛寫入的那筆資料
	var user User
	db.Where(User{PartnerID: 0, Name: ""})
	db.Find(&user)
	// 印出select 結果(正確用法)
	fmt.Print(user)

	// 這樣會印出一串記憶體位置，我不知道這是幹嘛，使用觀念好像跟以前有點不一樣，如上golang是丟了user進去，又將user拿來讀取
	// result := db.Find(&user)
	// fmt.Print(result)
}

//OfferEventJoinExcel 資料結構
type OfferEventJoinExcel struct {
	ID            int     `gorm:"primary_key;column:id"`
	HallID        int     `gorm:"column:hall_id"`
	Currency      string  `gorm:"column:currency"`
	Name          string  `gorm:"column:name"`
	StartDate     string  `gorm:"column:start_date"`
	EndDate       string  `gorm:"column:end_date"`
	ReceivedType  string  `gorm:"column:received_type"`   //領取方式(S:系統派發, R:玩家自領)
	IsAutoPayment string  `gorm:"column:is_auto_payment"` //自領到期後是否自動派發(Y/N)
	Complex       float64 `gorm:"column:complex"`
	AuditSwitch   string  `gorm:"column:audit_switch"` //是否需要稽核(Y/N)
	OpCode        int     `gorm:"column:opcode"`
	EventStatus   string  `gorm:"column:offer_event.status"` //活動狀態(1:未開始, 2:進行中, 3:已結束)
	EventID       int     `gorm:"column:event_id"`           //活動ID
	ExcelStatus   string  `gorm:"column:A"`                  //派發狀態(PN:派發未執行、PU:玩家自領中、PF:派發已處理、CR:沖銷處理中、CF:沖銷已完成)
	// 使用 `gorm:"column:offer_event_excel.status"` 會有被 `gorm:"column:offer_event.status"` 污染的情形發生，改使用別名

	AmountTotal float64 `gorm:"column:amount_total"` //總金額
	UserTotal   int     `gorm:"column:user_total"`   //總人數
}

func test2() {
	db, err := gorm.Open("mysql", "root:qwe123@tcp(127.0.0.1:3306)/AlmondDB?charset=utf8&parseTime=True&loc=Local")
	// 選擇要不要打開debug mode 超實用
	// LogMode set log mode, `true` for detailed logs, `false` for no log, default, will only print error logs
	db.LogMode(true)

	// 連不連得到，這個超級重要!!
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	var (
		HallID              = 6
		offerEventJoinExcel OfferEventJoinExcel
	)

	db.Table("offer_event").Select(
		`offer_event.id,
offer_event.hall_id,
offer_event.currency,
offer_event.name,
offer_event.start_date,
offer_event.end_date,
offer_event.received_type,
offer_event.is_auto_payment,
offer_event.complex,
offer_event.audit_switch,
offer_event.opcode,
offer_event.status,
offer_event_excel.id,
offer_event_excel.event_id,
offer_event_excel.status as A,
offer_event_excel.amount_total,
offer_event_excel.user_total`).Joins("inner join offer_event_excel on offer_event.id = offer_event_excel.event_id ").Where("offer_event.hall_id = ? ", HallID).Find(&offerEventJoinExcel)

	log.Println(offerEventJoinExcel)
}
