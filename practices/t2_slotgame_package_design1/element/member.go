package element

// Mem 單人遊戲運行數據
type Mem struct {
	// 免費遊戲次數
	FreeGameTimes int
	// 免費遊戲倍率
	FreeGameRatio int
	// 剩餘分數
	Credit int
}

// NewMemStruct 新增一個Mem結構
func NewMemStruct(credit int) *Mem {

	m := &Mem{
		FreeGameTimes: 0,
		FreeGameRatio: 1,
		Credit:        credit,
	}

	return m
}
