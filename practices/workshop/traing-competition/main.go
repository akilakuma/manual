package main

import (
	"log"
	"sync"
	"time"

	GOset "github.com/deckarep/golang-set"
)

// CreatorGod 創造之神
var CreatorGod *God

// God 創造之神型態
type God struct {
	liveList            []string             // 還存在的種族紀錄
	ancestorMap         map[string]*Ancestor // key:種族名稱, 先祖資料
	ancestorMapRWLocker *sync.RWMutex
	childMap            map[*Child]bool // key:後代個體
	childMapRWLocker    *sync.RWMutex
	winnerChan          chan struct{}
}

// periodFight 定期戰爭
func periodFight() {

	time.Sleep(1 * time.Second)

	aTicker := time.NewTicker(500 * time.Millisecond)
	cTicker := time.NewTicker(200 * time.Millisecond)

	for {
		select {
		case <-aTicker.C:
			log.Println()
			log.Println("發動先祖戰爭")
			go ancestorsFight()
		case <-cTicker.C:
			log.Println()
			log.Println("發動後代大亂鬥")
			go childenFight()
		case <-CreatorGod.winnerChan:
			log.Println()
			log.Println("遊戲結束")

			if len(CreatorGod.ancestorMap) > 0 {
				for kindName := range CreatorGod.ancestorMap {
					log.Println(kindName, "獲得勝利～")
					break
				}
			} else {
				log.Println("後代大亂鬥中同時陣亡，沒有種族獲得最後勝利")
			}
			return

		}
	}
}

func main() {

	CreatorGod = &God{
		ancestorMap:         make(map[string]*Ancestor), // key:種族名稱, 先祖資料
		ancestorMapRWLocker: new(sync.RWMutex),
		childMap:            make(map[*Child]bool), // key:後代個體
		childMapRWLocker:    new(sync.RWMutex),
		winnerChan:          make(chan struct{}),
	}

	// 創建10個種族
	nameList := []string{"Human", "Elf", "Orc", "Vampire", "Lizard", "Dragon", "Undead", "Dwarf", "Giant", "Werewolf"}
	for key, name := range nameList {
		go createAncestor(key, name)
	}

	CreatorGod.liveList = nameList

	periodFight()
}

// DifferSet 兩個[]string 取差集
func DifferSet(source []string, target []string) []string {
	sourceSet := GOset.NewSetFromSlice(StringToInterface(source))
	targetSet := GOset.NewSetFromSlice(StringToInterface(target))
	operatedSet := sourceSet.Difference(targetSet)
	result := operatedSet.ToSlice()
	return InterFaceToString(result)
}

// StringToInterface []String  轉換成  []interface{}
func StringToInterface(soruce []string) []interface{} {

	r := make([]interface{}, len(soruce))
	for i := range soruce {
		r[i] = interface{}(soruce[i])
	}
	return r
}

// InterFaceToString []interface{} 轉換成 []String
func InterFaceToString(soruce []interface{}) []string {

	r := make([]string, len(soruce))
	for i := range soruce {
		r[i] = soruce[i].(string)
	}
	return r
}
