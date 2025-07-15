package main

func QueryFirst(uId int64) DailyGame {

	var d DailyGame
	OrderDBStr.Table("daily_game").LogMode(true).Where("user_id = ? ", uId).Order("date desc").First(&d)
	return d
}

type DailyGame struct {
	Date       string  `gorm:"column:date"`         // 日期',
	GameName   string  `gorm:"column:game_name"`    // 遊戲名稱
	GameCnName string  `gorm:"column:game_cn_name"` // 遊戲中文名稱
	UserId     int64   `gorm:"column:user_id"`      // user id',
	UserName   string  `gorm:"column:user_name"`    // user 帳號名稱',
	BetGold    float64 `gorm:"column:bet_gold"`     // 總下注金額',
	BetNum     int     `gorm:"column:bet_num"`      // 注數',
	Profit     float64 `gorm:"column:profit"`       // 盈虧=派彩金額-投注金額(含退水)'
	Cnt        int     `gorm:"column:cnt"`          // 單數'
}
