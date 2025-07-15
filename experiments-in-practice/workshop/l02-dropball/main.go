package main

import (
	"log"
	"time"
)

// 當發現球無法傳遞給下一位時候，打電話通知

// method :
// Ａ action: 跟上帝接球，打電話，傳球給B
// B action: 跟A接球，打電話，接電話，傳球給C
// C action: 接球，處理球，接電話

// 訊息機器 :接收訊息，送訊息

// field ：
// A,B,C :
// a : channel 拿球，不限
// b : channel 拿球，2
// c : channel 拿球，2
// 訊息機器可以共用

type Ball struct {
	id int
}

type A struct {
	ballBasket chan Ball
	heaven     *Heaven
}

type B struct {
	ballBasket chan Ball
	heaven     *Heaven
}

type C struct {
	ballBasket chan Ball
	heaven     *Heaven
}

// Heaven 天堂
type Heaven struct {
	a *A
	b *B
	c *C
}

func main() {

	a := &A{
		ballBasket: make(chan Ball, 1000),
	}

	b := &B{
		ballBasket: make(chan Ball, 2),
	}

	c := &C{
		ballBasket: make(chan Ball, 2),
	}

	var (
		ball Ball
	)
	var heaven = &Heaven{
		a: a,
		b: b,
		c: c,
	}

	b.heaven = heaven
	c.heaven = heaven

	go heaven.b.getBall()
	go heaven.c.getBall()

	for i := 0; i < 10; i++ {
		ball = Ball{
			id: i,
		}
		a.passBall(b, ball)
	}

	time.Sleep(10 * time.Second)

}

func (a *A) passBall(b *B, ball Ball) {
	select {
	case b.ballBasket <- ball:
		// log.Println("球id", ball.id)
		log.Println("A success pass ball to B")
		time.Sleep(50 * time.Millisecond)
	default:
		// 打電話
		log.Println("call b")
	}
}

func (b *B) getBall() {

	for ball := range b.ballBasket {

		log.Println("B拿到球id", ball.id)

		// 球抽出來
	Loop:
		for {
			select {
			case b.heaven.c.ballBasket <- ball:
				log.Println("B success pass ball to C")
				break Loop
			default:
				log.Println("call c")
				time.Sleep(10 * time.Millisecond)
			}
		}
	}
}

func (c *C) getBall() {

	for ball := range c.ballBasket {
		time.Sleep(20 * time.Millisecond)
		log.Println("Ｃ把球丟掉", ball.id)
	}
}
