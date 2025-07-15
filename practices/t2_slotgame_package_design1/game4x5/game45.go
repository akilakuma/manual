package game4x5

import (
	Element "golang-advance-practice/t2_slotgame_package_design1/element"
	GameSetting "golang-advance-practice/t2_slotgame_package_design1/gamesetting"
	"math/rand"
	"strings"
	"time"
)

// Game45Body 4x5 的遊戲主體
type Game45Body struct {
	spinIndex  [5]int
	spinResult [5][4]string
}

// NewGameBody 建立一個遊戲處理組實體
func NewGameBody() GameSetting.GameBodyInterface {
	return &Game45Body{}
}

// Spin 產出挑選彩球結果
func (g *Game45Body) Spin(e *Element.Gset, kind string) {

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
func (g *Game45Body) SetNewSpinResultToCode(e *Element.Gset, kind string) {

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
			cIndex1 = codeIndex
			cIndex2 = codeIndex + 1
			cIndex3 = codeIndex + 2
			cIndex4 = codeIndex + 3
		)

		// 初始的位置在最後
		if codeIndex == len(splitResult) {
			cIndex2 = 0
			cIndex3 = 1
			cIndex4 = 2
		} else if codeIndex+1 == len(splitResult) {
			// 初始的位置+1會在最後
			cIndex3 = 0
			cIndex4 = 1
		} else if codeIndex+2 == len(splitResult) {
			// 初始的位置+2會在最後
			cIndex4 = 0
		}

		g.spinResult[i][0] = splitResult[cIndex1]
		g.spinResult[i][1] = splitResult[cIndex2]
		g.spinResult[i][2] = splitResult[cIndex3]
		g.spinResult[i][3] = splitResult[cIndex4]
	}
}

// CountMGCredit // 計算主要遊戲得分
func (g *Game45Body) CountMGCredit(mem *Element.Mem) {

}

// CountFGCredit // 計算免費遊戲得分
func (g *Game45Body) CountFGCredit(mem *Element.Mem) {

}
