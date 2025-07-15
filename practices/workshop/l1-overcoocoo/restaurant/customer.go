package restaurant

import (
	"strconv"
	"time"

	"github.com/fatih/color"
)

// Customer 客人
type Customer struct {
	id       int32  // 拿到的號碼牌(流水號)
	food     string // 想吃的餐點
	hasEaten bool   // 吃飽沒
}

// NewCustomer 新客人
func NewCustomer(food string) *Customer {
	return &Customer{
		food: food,
	}
}

// SetID 記錄編號
func (c *Customer) SetID(id int32) {
	c.id = id
}

// order 點餐
func (c *Customer) order() string {
	return c.food
}

// eating 吃飯
func (c *Customer) eating(food string) {
	color.Yellow(strconv.Itoa(int(c.id)) + "客人吃飯:" + food)
	time.Sleep(500 * time.Millisecond)
	c.leave()
}

// leave 閃人
func (c *Customer) leave() {
	c.hasEaten = false
	color.Yellow(strconv.Itoa(int(c.id)) + "客人吃飽離開")
}
