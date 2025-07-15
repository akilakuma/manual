package main

import (
	"context"
	"log"
	"math/rand"
	"time"
)

// Intermediary 仲介
type Intermediary struct {
	BadBadGuy
	TopBadGuy
	DeployRewardChan chan NormalContract
}

// MakeOtherFakeMission 製作假任務
func (i *Intermediary) MakeOtherFakeMission(fakenum int) []NormalContract {

	var fakeContracts []NormalContract
	for i := 0; i < fakenum; i++ {
		fakeContracts = append(fakeContracts, NormalContract{isfake: true})
	}
	return fakeContracts
}

// SendMissionToAssassion 將任務送給殺手們
func (i *Intermediary) SendMissionToAssassion(missions []NormalContract) {

	for _, mission := range missions {
		i.DeployRewardChan <- mission
	}
}

// ReWardMurderer 懸賞兇手
func (i *Intermediary) ReWardMurderer(name string) {
	for killerName, killerChan := range i.rewardMap {
		if killerName != name {
			killerChan <- EmergencyMsg{
				killerName: name,
			}
		}
	}
}

// Action 行動
func (i *Intermediary) Action(ctx context.Context) {

loop:
	for {
		select {
		// 被結束生命了
		case <-ctx.Done():

			ctxData := ctx.Value("nameData").(*CtxValue)
			killerName := ctxData.GetName()
			i.emergencyNotifyChan <- EmergencyMsg{
				killerName: killerName,
			}
			break loop
		// 收到一般任務
		case realMission, ok := <-i.normalContractChan:

			if ok {
				log.Println("仲介收到任務")

				// 製作假情報任務
				fakeMissions := i.MakeOtherFakeMission(29)

				// 將假的任務混在真的裡面
				rand.Seed(time.Now().UnixNano())
				x := rand.Intn(len(fakeMissions))
				missions := insertToSlice(x, realMission, fakeMissions)

				// 丟給殺手們
				i.SendMissionToAssassion(missions)
			} else {
				break loop
			}
			// 收到緊急訊息
		case emergencyMsg, ok := <-i.emergencyNotifyChan:

			if ok {
				log.Println("仲介收到緊急訊息:", emergencyMsg)

				i.ReWardMurderer(emergencyMsg.killerName)

			}

			break loop
		}
	}

}
