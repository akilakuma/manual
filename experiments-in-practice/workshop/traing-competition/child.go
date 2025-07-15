package main

import (
	"context"
	"log"
	"math/rand"
	"sync/atomic"
	"time"
)

// Child 後代型態
type Child struct {
	ancestor *Ancestor          // 先祖
	id       int32              // 編號
	kindName string             // 種族名稱
	ctx      context.Context    // 概念比較像是『靈魂』？
	ctxDead  context.CancelFunc // 死亡method
}

// 創造1~3個後代
func createChildren(parentCtx context.Context, kindName string, a *Ancestor) {

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(2) + 1
	// num := 1

	for i := 0; i < num; i++ {
		childCtx, childDead := context.WithCancel(parentCtx)

		c := &Child{
			ancestor: a,
			id:       atomic.AddInt32(&a.childID, 1),
			ctx:      childCtx,
			ctxDead:  childDead,
			kindName: kindName,
		}
		recordChild(c)
		go c.Action()
	}

}

// Action 後代行動
func (c *Child) Action() {

	t := time.NewTimer(100 * time.Millisecond)

	for {
		select {
		case <-t.C:
			createChildren(c.ctx, c.kindName, c.ancestor)

		case <-c.ctx.Done():
			// 告訴下面的也死一死吧

			CreatorGod.childMapRWLocker.Lock()
			delete(CreatorGod.childMap, c)
			CreatorGod.childMapRWLocker.Unlock()

			c.ctxDead()
			return

		}
	}
}

// 出生的後代，註冊在案
func recordChild(c *Child) {
	CreatorGod.childMapRWLocker.Lock()
	CreatorGod.childMap[c] = false
	CreatorGod.childMapRWLocker.Unlock()
}

// 後代大亂鬥
func childenFight() {

	num := 150
	deleteChild(num)
}

// 隨機挑選X個後代死亡
func deleteChild(num int) {
	CreatorGod.childMapRWLocker.Lock()
	var i = 1

	for child := range CreatorGod.childMap {

		delete(CreatorGod.childMap, child)
		child.ctxDead()
		i++
		if i >= num {
			break
		}
	}
	CreatorGod.childMapRWLocker.Unlock()

	checkAllChildrenNum()
}

// 檢查是否所有後代已陣亡
func checkAllChildrenNum() {

	var statisticMap = make(map[string]int)

	CreatorGod.childMapRWLocker.Lock()
	for child := range CreatorGod.childMap {
		statisticMap[child.kindName]++
	}
	CreatorGod.childMapRWLocker.Unlock()

	var liveKindList []string
	for kindName, num := range statisticMap {
		log.Println("種族:", kindName, ",剩餘後代數量:", num)
		liveKindList = append(liveKindList, kindName)
	}

	// 有可以一個種族以上一同滅絕
	CreatorGod.ancestorMapRWLocker.Lock()

	extinctList := DifferSet(CreatorGod.liveList, liveKindList)

	CreatorGod.ancestorMapRWLocker.Unlock()

	if len(extinctList) > 0 {
		log.Println("🐥🐥🐥因後代全數陣亡，先祖連帶敗亡:", extinctList)
		killAncestors(extinctList)
	}

	if len(CreatorGod.ancestorMap) <= 1 {
		CreatorGod.winnerChan <- struct{}{}
	}

}
