package main

import (
	"log"
	"time"
)

/*

- 問題：
	如果close channel，channel的接收端會不會自動拿到！ok

-結論：
	會，不需要額外往channel塞什麼資料作通知，for迴圈抽channel時就知道已被關閉。

-output：
	2025/07/15 13:50:39 收到貸款 loan: 欠錢要還 50000
	2025/07/15 13:50:39 收到貸款 loan: 利上滾利 100000
	2025/07/15 13:50:39 channel 關閉，離開receive

*/

func practice1() {

	var bank = Bank{
		name:        "YY bank",
		contactChan: make(chan *LendContact, 10),
	}

	go bank.receiveLoan()

	bank.contactChan <- &LendContact{
		msg:   "欠錢要還",
		money: 50000,
	}

	bank.contactChan <- &LendContact{
		msg:   "利上滾利",
		money: 100000,
	}

	// 關閉channel
	close(bank.contactChan)

	for {
		time.Sleep(100 * time.Second)
	}
}

type (
	// Bank 銀行
	Bank struct {
		name        string
		contactChan chan *LendContact
	}

	// LendContact 借貸通道
	LendContact struct {
		msg   string
		money int
	}
)

func (b *Bank) receiveLoan() {
	for {
		loan, ok := <-b.contactChan
		if ok {
			log.Println("收到貸款 loan:", loan.msg, loan.money)
		} else {
			// b.contactChan 已經關閉，會執行到這裡
			break
		}
	}

	log.Println("channel 關閉，離開receive")
}
