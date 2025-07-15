package main

import (
	"context"
	"log"
	"time"
)

// Boss 幕後老闆
type Boss struct {
	BadBadGuy
	TopBadGuy
}

// MakeMission 製作任務
func (b *Boss) MakeMission() NormalContract {
	return NormalContract{}
}

// SendMissionToIntermediary 將任務丟給仲介
func (b *Boss) SendMissionToIntermediary(n NormalContract) {
	b.normalContractChan <- n
}

// ReWardMurderer 懸賞兇手
func (b *Boss) ReWardMurderer(name string) {
	for killerName, killerChan := range b.rewardMap {
		if killerName != name {
			killerChan <- EmergencyMsg{
				killerName: name,
			}
		}
	}
}

// Action 行動
func (b *Boss) Action(ctx context.Context) {

	// 每50秒觸發一次
	t := time.NewTicker(50 * time.Millisecond)

loop:
	for {
		select {
		// 被結束生命了
		case <-ctx.Done():

			ctxData := ctx.Value("nameData").(*CtxValue)
			killerName := ctxData.GetName()
			b.emergencyNotifyChan <- EmergencyMsg{
				killerName: killerName,
			}
			break loop

		// 發佈新任務給仲介
		case <-t.C:

			log.Println("boss 發布任務")
			contract := b.MakeMission()
			b.SendMissionToIntermediary(contract)

		// 收到緊急訊息
		case emergencyMsg, ok := <-b.emergencyNotifyChan:

			if ok {
				log.Println("老闆收到緊急訊息:", emergencyMsg)

				b.ReWardMurderer(emergencyMsg.killerName)

			}
			break loop
		}
	}
}
