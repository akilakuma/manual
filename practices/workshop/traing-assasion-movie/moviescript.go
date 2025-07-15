package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

// MovieScript 電影劇本管理
type MovieScript struct {
	bossCtx            context.Context
	bossCancel         context.CancelFunc
	intermediaryCtx    context.Context
	intermediaryCancel context.CancelFunc
	killerActorMap     map[*Assassin]bool
	killerRWLocker     *sync.RWMutex
	IntermediaryActor  *Intermediary
	bossActor          *Boss
}

var movieScript = MovieScript{
	killerActorMap: make(map[*Assassin]bool, 0),
	killerRWLocker: new(sync.RWMutex),
}

// KillHighLevelPerson 幹掉上層
func (m *MovieScript) KillHighLevelPerson(name string) {

	// 隨機挑選仲介或幕後老闆
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(2)

	if x == 0 {
		// 老闆掰
		data := m.bossCtx.Value("nameData").(*CtxValue)
		data.SetName(name)
		log.Println("老闆掰")
		m.bossCancel()
	} else {
		// 仲介掰
		data := m.intermediaryCtx.Value("nameData").(*CtxValue)
		data.SetName(name)
		log.Println("仲介掰")
		m.intermediaryCancel()
	}
}

// KillMurderer 幹掉兇手
func (m *MovieScript) KillMurderer(name, revengeName string) {
	for killer := range m.killerActorMap {
		m.killerRWLocker.Lock()
		if killer.name == name && killer.isLive {
			killer.isLive = false
			log.Println(revengeName, "殺了:", killer.name)
			m.Close()
		}
		m.killerRWLocker.Unlock()
	}
}

// Close 關閉必要channel
func (m *MovieScript) Close() {

	// 關了仲介和老闆的channel
	// 就會break仲介和老闆的goroutine
	close(m.IntermediaryActor.normalContractChan)

	// 關了仲介和殺手們的channel
	// 就會break殺手們的goroutine
	close(m.IntermediaryActor.DeployRewardChan)
}

// CtxValue context 內容
type CtxValue struct {
	name string
}

// SetName 設定名字
func (c *CtxValue) SetName(name string) {
	c.name = name
}

// GetName 取得名字
func (c *CtxValue) GetName() string {
	return c.name
}
