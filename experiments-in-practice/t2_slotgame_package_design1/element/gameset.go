package element

// Gset 遊戲設定
type Gset struct {
	// 遊戲類型
	Category string
	// 遊戲ID
	ID int
	// 主要遊戲彩球(5條)
	MainGameCode [5]string
	// 免費遊戲彩球(5條)
	FreeGameCode [5]string
}

// NewGameSettingStruct 新增一個Gset結構
func NewGameSettingStruct(category string, ID int, mainGameCode, freeGameCode [5]string) *Gset {

	g := &Gset{
		Category:     category,
		ID:           ID,
		MainGameCode: mainGameCode,
		FreeGameCode: freeGameCode,
	}

	return g
}
