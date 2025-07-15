package main

import (
	"context"
	"log"
	"runtime"
	"strconv"
	"time"
)

// BadBadGuy 幹壞事的傢伙
type BadBadGuy struct {
	emergencyNotifyChan chan EmergencyMsg    // 緊急聯絡管道
	normalContractChan  chan *NormalContract // 一般任務承接管道
}

// TopBadGuy 最壞的那些傢伙
type TopBadGuy struct {
	rewardMap map[string](chan EmergencyMsg)
}

// EmergencyMsg 緊急通知
type EmergencyMsg struct {
	killerName string
}

// NormalContract 一般訊息
type NormalContract struct {
	id     int
	isfake bool
}

func main() {

	// 背後老闆和仲介的任務溝通管道
	b2Inormal := make(chan *NormalContract, 1000)
	b2Iemergency := make(chan EmergencyMsg)

	// 建立幕後老闆角色
	BlackBoss := createBoss(b2Inormal, b2Iemergency)

	// 建立仲介角色
	BlackIntermediary := createIntermediary(b2Inormal, b2Iemergency)

	// 僱傭殺手們
	// 仲介提供發布任務給殺手們的訊息管道
	John, JohnEmergency := hireKiller("John Wick", BlackIntermediary.DeployRewardChan)
	Jason, JasonEmergency := hireKiller("Jason Bourne", BlackIntermediary.DeployRewardChan)
	Ethan, EthanEmergency := hireKiller("Ethan Hunt", BlackIntermediary.DeployRewardChan)

	go John.Action()
	go Jason.Action()
	go Ethan.Action()

	// 老闆和仲介跟殺手們各自約定的緊急聯絡方式
	killerEmergncy := map[string](chan EmergencyMsg){
		"John Wick":    JohnEmergency,
		"Jason Bourne": JasonEmergency,
		"Ethan Hunt":   EthanEmergency,
	}
	BlackBoss.rewardMap = killerEmergncy
	BlackIntermediary.rewardMap = killerEmergncy

	log.Println("🐥🐥 一開始的 gorutine 數量:" + strconv.Itoa(runtime.NumGoroutine()))
	t := time.NewTicker(3 * time.Second)
	for {
		<-t.C
		log.Println("🐥🐥 目前gorutine 數量:" + strconv.Itoa(runtime.NumGoroutine()))
		if runtime.NumGoroutine() == 1 {
			break
		}
	}
}

// 建立幕後老闆角色
func createBoss(b2Inormal chan *NormalContract, b2Iemergency chan EmergencyMsg) *Boss {

	BlackBoss := &Boss{
		BadBadGuy: BadBadGuy{
			normalContractChan:  b2Inormal,
			emergencyNotifyChan: b2Iemergency,
		},
	}

	bossCtxWithValue := context.WithValue(context.Background(), "nameData", &CtxValue{})
	bossCtx, bossCancel := context.WithCancel(bossCtxWithValue)

	movieScript.bossCtx = bossCtxWithValue
	movieScript.bossCancel = bossCancel
	movieScript.bossActor = BlackBoss

	go BlackBoss.Action(bossCtx)

	return BlackBoss
}

// 建立仲介角色
func createIntermediary(b2Inormal chan *NormalContract, b2Iemergency chan EmergencyMsg) *Intermediary {
	BlackIntermediary := &Intermediary{
		BadBadGuy: BadBadGuy{
			normalContractChan:  b2Inormal,
			emergencyNotifyChan: b2Iemergency,
		},
		DeployRewardChan: make(chan *NormalContract, 1000),
	}

	intermediaryCtxWithValue := context.WithValue(context.Background(), "nameData", &CtxValue{})
	intermediaryCtx, intermediaryCancel := context.WithCancel(intermediaryCtxWithValue)

	movieScript.intermediaryCtx = intermediaryCtxWithValue
	movieScript.intermediaryCancel = intermediaryCancel
	movieScript.IntermediaryActor = BlackIntermediary

	go BlackIntermediary.Action(intermediaryCtx)

	return BlackIntermediary
}

// 建立殺手角色
func hireKiller(name string, publicNormalContractChan chan *NormalContract) (*Assassin, chan EmergencyMsg) {
	emergency := make(chan EmergencyMsg)
	a := &Assassin{
		BadBadGuy: BadBadGuy{
			normalContractChan:  publicNormalContractChan,
			emergencyNotifyChan: emergency,
		},
		name:   name,
		isLive: true,
	}

	movieScript.killerActorMap[a] = true

	return a, emergency
}

// 插入到slice裡面
func insertToSlice(index int, para NormalContract, a []NormalContract) []NormalContract {

	split1 := make([]NormalContract, len(a[:index]))
	copy(split1, a[:index])
	split2 := make([]NormalContract, len(a[index:]))
	copy(split2, a[index:])

	var rtnSlice []NormalContract
	rtnSlice = append(split1, para)
	rtnSlice = append(rtnSlice, split2...)

	return rtnSlice
}
