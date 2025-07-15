package game3x5

import (
	Element "golang-advance-practice/t2_slotgame_package_design1/element"
	GameSetting "golang-advance-practice/t2_slotgame_package_design1/gamesetting"
	"math/rand"
	"strings"
	"time"
)

// Game35Body 3x5 的遊戲主體
type Game35Body struct {
	spinIndex  [5]int
	spinResult [5][3]string
}

// NewGameBody 建立一個遊戲處理組實體
func NewGameBody() GameSetting.GameBodyInterface {
	return &Game35Body{}
}

// Spin 產出挑選彩球結果
func (g *Game35Body) Spin(e *Element.Gset, kind string) {

	// 亂數產生
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// 結果寫回array
	for i := 0; i < len(g.spinIndex); i++ {
		g.spinIndex[i] = r1.Intn(100)
	}

	g.SetNewSpinResultToCode(e, kind)
}

// SetNewSpinResultToCode 亂數mapping到彩球
func (g *Game35Body) SetNewSpinResultToCode(e *Element.Gset, kind string) {

	// c := [1][2]int{{1,2}}

	var soureGameCode [5]string

	switch kind {
	case "main":
		soureGameCode = e.MainGameCode
	case "free":
		soureGameCode = e.FreeGameCode
	case "default":
		soureGameCode = e.MainGameCode
	}

	for i := 0; i < len(g.spinIndex); i++ {
		splitResult := strings.Split(soureGameCode[i], ",")

		// ex : 100 % 6 = 4, 挑index = 4 放在中間
		codeIndex := g.spinIndex[i] % len(splitResult)

		var (
			cIndex1 = codeIndex - 1
			cIndex2 = codeIndex
			cIndex3 = codeIndex + 1
		)

		// 中間的位置在0，前一個會變成最後一個
		if codeIndex == 0 {
			cIndex1 = len(splitResult)
		} else if codeIndex == len(splitResult) {
			// 中間的位置在最後一個，後一個會變成是第一個
			cIndex3 = 0
		}

		g.spinResult[i][0] = splitResult[cIndex1]
		g.spinResult[i][1] = splitResult[cIndex2]
		g.spinResult[i][2] = splitResult[cIndex3]
	}
}

// CountMGCredit // 計算主要遊戲得分
func (g *Game35Body) CountMGCredit(mem *Element.Mem) {

}

// CountFGCredit // 計算免費遊戲得分
func (g *Game35Body) CountFGCredit(mem *Element.Mem) {

}
