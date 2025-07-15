package main

import (
	"log"
	"time"
)

// Assassin 殺手
type Assassin struct {
	BadBadGuy
	name      string // 殺手名字
	executime int    // 執行真的任務次數
	isLive    bool
}

// DetectMission 判斷任務真假
func (a *Assassin) DetectMission(con NormalContract) {

	if con.isfake {
		a.AbortFakeMission()
	} else {
		a.ExecuteRealMission()
	}
}

// AbortFakeMission 丟棄假的任務
func (a *Assassin) AbortFakeMission() {
	time.Sleep(1 * time.Millisecond)
}

// ExecuteRealMission 執行真的任務
func (a *Assassin) ExecuteRealMission() {
	time.Sleep(10 * time.Millisecond)
	a.executime++
	log.Println("殺手:", a.name, "執行到真的任務,次數:", a.executime)
}

// Action 殺手行動
func (a *Assassin) Action() {
loop:
	for {
		select {
		case mission, ok := <-a.normalContractChan:

			if ok {
				a.DetectMission(mission)

				if a.executime >= 20 {
					movieScript.KillHighLevelPerson(a.name)
					break loop
				}
			} else {
				break loop
			}

		case emergencyMsg, ok := <-a.emergencyNotifyChan:
			if ok {
				log.Println("殺手:", a.name, ",收到緊急任務:", emergencyMsg)
				movieScript.KillMurderer(emergencyMsg.killerName, a.name)
			} else {

				break loop
			}
		}
	}
}
