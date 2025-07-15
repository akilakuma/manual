package restaurant

import (
	"time"

	"github.com/fatih/color"
)

// Chief 廚師
type Chief struct {
	job       chan string
	onDuty    chan bool
	foodTable chan string
}

// NewChief 新廚師
func NewChief() *Chief {
	c := &Chief{
		job:    make(chan string, 100),
		onDuty: make(chan bool),
	}

	go c.Ｗorking()

	return c
}

// setOutputFoodTable 告訴廚師出餐的位置
func (c *Chief) setOutputFoodTable(table chan string) {
	c.foodTable = table
}

// getOrder 取得訂單
func (c *Chief) getOrder(jobs <-chan string) {
	for food := range jobs {
		c.job <- food
	}
}

// cook 做菜
func (c *Chief) cook(food string) {

	color.Green("廚師製作餐點:" + food)
	// 假設這是做菜
	time.Sleep(100 * time.Millisecond)

	// 出餐
	c.sendOut(food)
}

// sendOut 出菜
func (c *Chief) sendOut(food string) {
	color.Green("廚師出菜:" + food)
	c.foodTable <- food
}

// takeBreak 下班休息
func (c *Chief) takeBreak() {
	c.onDuty <- false
}

// Ｗorking 工作
func (c *Chief) Ｗorking() {

Loop:
	for {
		select {
		case food, ok := <-c.job:
			if ok {
				c.cook(food)
			}
		case status, ok := <-c.onDuty:
			if ok {
				if !status {
					break Loop
				}
			}
		}
	}
}
