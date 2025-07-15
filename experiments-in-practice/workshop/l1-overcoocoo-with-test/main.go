package main

import (
	"golang-advance-practice/workshop/l1-overcoocoo-with-test/restaurant"
	"log"
	"time"
)

func main() {

	var city City
	city.EatingTrip()

	time.Sleep(100 * time.Second)
	log.Println("程式結束")
}

// City 城市
type City struct {
	r *restaurant.Restaurant
}

// EatingTrip 吃飯之旅開始
func (c *City) EatingTrip() {
	// 建立餐廳
	c.r = restaurant.CreateRestaurant()

	time.Sleep(1 * time.Second)

	for i := 0; i < 30; i++ {

		// 新客人
		log.Println("新客人 coming")

		fookind := i % 5
		var ordrFood string
		switch fookind {
		case 0, 1:
			ordrFood = "salar"
		case 2, 3:
			ordrFood = "pizza"
		case 4:
			ordrFood = "steak"
		}
		customer := restaurant.NewCustomer(ordrFood)

		// 給客人一個號碼牌(流水號)
		cID := c.r.NewCustomerID()
		customer.SetID(cID)

		// 進入餐廳
		log.Println("新客人:", cID, "進入餐廳")
		err := c.r.CustomerComein(customer)
		if err != nil {
			log.Println("因為", err, "新客人:", cID, "森77離開了")
		}

		time.Sleep(50 * time.Millisecond)
	}

}
