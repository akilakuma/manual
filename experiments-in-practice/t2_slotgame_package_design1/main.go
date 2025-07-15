package main

import (
	Element "golang-advance-practice/t2_slotgame_package_design1/element"
	Game3x5 "golang-advance-practice/t2_slotgame_package_design1/game3x5"
	Game4x5 "golang-advance-practice/t2_slotgame_package_design1/game4x5"
	GameSetting "golang-advance-practice/t2_slotgame_package_design1/gamesetting"
)

var isMain = true

func main() {

	// 可以改由其他來源傳入餘額
	var personalCredit = 1000
	// 建立一個新的使用者遊戲連線資訊
	mem := Element.NewMemStruct(personalCredit)

	// 取得遊戲設定資訊
	// 假設從DB讀取設定
	category, ID, mainGameCode, freeGameCode := getDBSetting(1)
	// 取得設定
	gSet := Element.NewGameSettingStruct(category, ID, mainGameCode, freeGameCode)
	gameBody := getNewGameBody(gSet)

	// 進行遊戲
	for {

		if isMain {

			// Main Game Spining 拉霸一次
			gameBody.Spin(gSet, "main")

			// Count Main Game Credit 計算Main Game 得分
			gameBody.CountMGCredit(mem)

			// Whether Get Free Game
			GameSetting.AddFreeGame(mem, 1)
			GameSetting.SetFreeGameRatio(mem, 3)

		} else {

			// Free Game Spining 拉霸一次
			gameBody.Spin(gSet, "free")

			// Count Free Game Credit 計算Free Game 得分
			gameBody.CountFGCredit(mem)
		}

	}

}

// getDBSetting 從DB讀取設定，假的，意思到了就好
func getDBSetting(ID int) (string, int, [5]string, [5]string) {
	switch ID {
	case 1:
		return "3x5", 1, [5]string{"2,3,4,5,6", "3,2,4,3,4", "1,2,1,3,4", "2,2,4,2,4", "4,2,4,3,5"}, [5]string{"1,1,4,5,6", "2,2,4,3,2", "1,2,1,3,4", "2,2,4,2,5", "5,2,4,3,5"}
	default:
		return "4x5", 1, [5]string{"2,3,4,5,6", "3,2,4,3,4", "1,2,1,3,4", "2,2,4,2,4", "4,2,4,3,5"}, [5]string{"1,1,4,5,6", "2,2,4,3,2", "1,2,1,3,4", "2,2,4,2,5", "5,2,4,3,5"}
	}

}

// getNewGameBody 取得gamebody
func getNewGameBody(gset *Element.Gset) GameSetting.GameBodyInterface {
	switch gset.Category {
	case "3x5":
		return Game3x5.NewGameBody()

	case "4x5":
		return Game4x5.NewGameBody()
	default:

		return nil
	}

}
