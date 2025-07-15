package main

import (
	"log"
	"time"
)

func main() {

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

	close(bank.contactChan)

	for {
		time.Sleep(100 * time.Second)
	}
}

// Bank 銀行
type Bank struct {
	name        string
	contactChan chan *LendContact
}

// LendContact 借貸通道
type LendContact struct {
	msg   string
	money int
}

func (b *Bank) receiveLoan() {
	for {
		loan, ok := <-b.contactChan
		if ok {
			log.Println("收到貸款 loan:", loan.msg, loan.money)
		} else {
			break
		}
	}

	log.Println("channel 關閉，離開receive")
}
