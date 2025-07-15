package main

import (
	"context"
	"log"
	"time"
)

// Ancestor 祖先型態
type Ancestor struct {
	childID  int32              // 子孫數量
	kindID   int                // 種族編號
	kindName string             // 種族名稱
	ctx      context.Context    // 概念比較像是『靈魂』？
	ctxDead  context.CancelFunc // 死亡method
}

// 創造先祖
func createAncestor(kindID int, kindName string) {

	ctx, ctxDead := context.WithCancel(context.Background())

	a := &Ancestor{
		kindID:   kindID,
		kindName: kindName,
		ctx:      ctx,
		ctxDead:  ctxDead,
	}
	CreatorGod.ancestorMapRWLocker.Lock()
	CreatorGod.ancestorMap[kindName] = a
	CreatorGod.ancestorMapRWLocker.Unlock()

	t := time.NewTimer(100 * time.Millisecond)
	<-t.C
	createChildren(a.ctx, a.kindName, a)
	select {
	case <-ctx.Done():
		// 整族滅光光
		ctxDead()
		return

	}
}

// ancestorsFight 先祖戰爭
func ancestorsFight() {

	var loser string
	CreatorGod.ancestorMapRWLocker.Lock()
	for _, ancestor := range CreatorGod.ancestorMap {
		loser = ancestor.kindName
		break
	}
	CreatorGod.ancestorMapRWLocker.Unlock()

	log.Println("先祖殞落:", loser)

	killAncestors([]string{loser})

	if len(CreatorGod.ancestorMap) <= 1 {
		CreatorGod.winnerChan <- struct{}{}
	}
}

// killAncestors 先祖死亡
func killAncestors(ancestors []string) {
	log.Println("先祖死亡名單:", ancestors)

	CreatorGod.ancestorMapRWLocker.Lock()
	for _, kindName := range ancestors {
		if ancestor, exists := CreatorGod.ancestorMap[kindName]; exists {
			ancestor.ctxDead()
			delete(CreatorGod.ancestorMap, kindName)
		}
	}

	deleteGodLiveList(ancestors)
	CreatorGod.ancestorMapRWLocker.Unlock()
}

// deleteGodLiveList 在創造之神的先祖名單中刪除
func deleteGodLiveList(name []string) {
	var newLiveSlice []string
	for _, n := range name {
		for index, nn := range CreatorGod.liveList {
			if n == nn {
				log.Println(n, "滅亡")
				newLiveSlice = append(CreatorGod.liveList[:index], CreatorGod.liveList[index+1:]...)
			}
		}
		CreatorGod.liveList = newLiveSlice
	}
}
