package gamesetting

import (
	Element "golang-advance-practice/t2_slotgame_package_design1/element"
)

// MGCreditCounter 計算主要遊戲得分-介面
type MGCreditCounter interface {
	CountMGCredit(mem *Element.Mem)
}

// FGCreditCounter 計算免費遊戲得分-介面
type FGCreditCounter interface {
	CountFGCredit(mem *Element.Mem)
}

// AddFreeGame 新增免費遊戲次數
func AddFreeGame(mem *Element.Mem, addTimes int) {
	mem.FreeGameTimes = mem.FreeGameTimes + 1
}

// SetFreeGameRatio 設定免費遊戲賠率
func SetFreeGameRatio(mem *Element.Mem, ratio int) {
	mem.FreeGameRatio = mem.FreeGameRatio + 1
}

// DefaultGameHandler 預設的遊戲處理組
// var DefaultGameHandler = &GameHandler{}

// GameHandler 遊戲處理組
// type GameHandler struct {
// 	Gamebody GameBodyInterface
// }

// GameBodyInterface GameBody的介面
type GameBodyInterface interface {
	Spin(e *Element.Gset, kind string)
	MGCreditCounter
	FGCreditCounter
}
