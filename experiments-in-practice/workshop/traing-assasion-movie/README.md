Golang 程式設計題目(殺手電影情節)
===

###### tags: `訓練`

## 練習目的
1. 練習goroutine與channel的搭配。
2. 練習channel的讀/寫/關閉使用，channel/bufferchannel。
3. 練習收拾開出去的goroutine與channel。

<br/>

## 題目說明

請用程式完成故事。

### 第一部分

John Wick、Jason Bourne、Ethan Hunt 三位是受雇的殺手，共同接受一位仲介傳過來的合約執行任務，而仲介的背後有一位幕後老闆。

幕後老闆會固定每經過50ms，給仲介一個任務，每次仲介收到後，為了安全性，會製造另外29假情報任務，所以每輪從仲介丟出的新任務出來，總共會有30個任務，其中29個是假的，只有一個是真實的任務。

John Wick、Jason Bourne、Ethan Hunt，拿到任務之後，都有辨識真偽的能力，判斷出是假的會丟棄不管，若是真的則會執行。

John Wick、Jason Bourne、Ethan Hunt，每次30個真偽任務出來都有公平的競爭機會去搶接，每次只能拿到一個，拿到之後，需要判斷真偽，假的做丟棄，真的則去執行任務並完成之後，才能去承接下一個。

判斷是假情報任務，需要花費1ms，執行真的任務，需要花費10ms。

一個任務，不論真假，不能同時被兩個人以上拿到，一個人一次只能拿一個，拿過就沒了。

### 第二部分

當三個殺手其中有人率先執行過20個真正任務的時候，他就想要養老退休了。

因為執行任務的夠多，發現了拿到的報酬與工作內容不對等的狀況，有種被欺騙的感覺，所以退休前生氣氣，要把上面的人幹掉。

仲介或幕後老闆因此隨機被幹掉了一個。

仲介和幕後老闆生前有個約定，遇到緊急狀況，在臨死之前會通知對方。

沒死的那位，當意外發生時，會收到對方唯一一次的通知，知道發生了事件，並且知道行兇的對象。

沒死的那位，會用緊急通知的方式，發佈一個刺殺任務，給另外兩位，兩位殺手都會收到，幹掉反叛的殺手可以獲得1000萬美金。

另外兩位因此會去執行任務，但只有一個人能夠先殺到反叛殺手，並且得到獎金。

### 第三部分

於是故事就落幕了，原本的角色就地解散，請關閉任何開出的goroutine或者channel。

<br/>

## 重點要求

本實作特別注重題目完成度與合理性，請務必讓實作內容符合題目劇本敘述。

儘量避免不合理的狀況，例如說當幕後老闆或仲介已經死了，殺手們還一直持續在收一般任務訊息。

例如同時有兩個人以上同時完成了20次任務執行(還沒做完的不算數)。

例如能夠一邊執行任務，一邊接受新任務。

等等類似不合理的狀況，請練習思考並模擬。

<br/>

## 範例

請以以下的code做改寫，自行調整或新增struct內容，和新增所需要的method或struct。


``` go
package main

// Assassin 殺手
type Assassin struct {
	emergencyNotifyChan chan *emergencyMsg
	normalContractChan  chan *normalContract
}

// Intermediary 仲介
type Intermediary struct{}

// Boss 幕後老闆
type Boss struct{}

// emergencyMsg 緊急通知
type emergencyMsg struct {
}

// normalContract 一般訊息
type normalContract struct {
}

// KillHighLevelPerson 幹掉上層
func KillHighLevelPerson() {

}

```
