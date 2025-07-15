package restaurant

import (
	"strconv"
	"time"

	"github.com/fatih/color"
)

// 服務生工作
// 優先檢查是否有剛入座的客人，為他們點餐
// 點餐後向廚師下單
// 檢查出菜台是否需要送菜

// Waiter 服務生
type Waiter struct {
	rt *Restaurant // 餐廳
}

// NewWaiter 新服務生
func NewWaiter(rt *Restaurant) *Waiter {
	w := &Waiter{
		rt: rt,
	}
	go w.working()
	return w
}

// serveOfCustomerOrder 為客人點餐服務
func (w *Waiter) serveOfCustomerOrder() {
	w.rt.seatRWlocker.Lock()

	for seatIndex, seatDetail := range w.rt.seats {
		// 有入座，但未點餐
		if seatDetail.food == "" {

			// 更新座位資訊
			w.rt.seats[seatIndex] = Seat{
				cID:      seatDetail.cID,
				customer: seatDetail.customer,
				food:     seatDetail.customer.food,
			}

			color.Cyan("服務生幫客人點餐:"+ strconv.Itoa(int(seatDetail.cID)))
			// log.Println("服務生幫客人點餐:", seatDetail.cID)
			// 獲得客人想點的餐點
			w.serveOfOrderToChief(seatIndex, seatDetail.cID, seatDetail.customer.food)
		}
	}
	w.rt.seatRWlocker.Unlock()
}

// serveOfOrderToChief 向廚師下單
func (w *Waiter) serveOfOrderToChief(seat int, cID int32, food string) {
	color.Cyan("服務生下單餐點:"+food+"是客人"+strconv.Itoa(int(cID))+ "的餐點")
	// 下完單就可以閃人，不用傻站到廚師煮完
	go w.rt.orderLine.income(seat, cID, food)
}

// serveOfDeliverFood 從廚師手中接菜送到客人手上
func (w *Waiter) serveOfDeliverFood() {
	// 檢查出餐檯
	select {
	case food, ok := <-w.rt.orderLine.outputFoodTable:
		if ok {

			// 從餐台的mark知道食物要送到哪個客人
			seatIndex := w.checkFoodForCustomer(food)
			// 知道是哪位客人
			customer := w.rt.seats[seatIndex].customer

			color.Cyan("服務生下單餐點:"+food+"座位:"+ strconv.Itoa(seatIndex)+ "給客人:"+ strconv.Itoa(int(w.rt.seats[seatIndex].customer.id)))
			// 拿餐點給客人吃飯
			customer.eating(food)
		}

	default:
		// 沒有餐點出餐
	}
}

// checkFoodForCustomer 回傳是哪張單，哪個座位的餐點
func (w *Waiter) checkFoodForCustomer(food string) int {

	w.rt.orderLine.markRWLocker.Lock()

	markSlice := w.rt.orderLine.markMap[food]
	seatIndex := markSlice[0].seatIndex

	// 把順位第一個點餐單移除
	if len(markSlice) != 1 {
		w.rt.orderLine.markMap[food] = append([]FoodMark{}, markSlice[1:len(markSlice)]...)
	} else {
		w.rt.orderLine.markMap[food] = append([]FoodMark{}, markSlice[0])
	}
	w.rt.orderLine.markRWLocker.Unlock()
	return seatIndex
}

// checkCustomerLeave 確認是否有客人吃飽離開
func (w *Waiter) checkCustomerLeave() {
	w.rt.seatRWlocker.Lock()
	for seatIndex, seatDetail := range w.rt.seats {
		// 入座的客人，已經點餐
		if seatDetail.food != "" {
			// 若吃飽了，清空座位
			if seatDetail.customer.hasEaten {
				delete(w.rt.seats, seatIndex)
			}
		}
	}
	w.rt.seatRWlocker.Unlock()
}

// working
func (w *Waiter) working() {

	for {

		// 確認是否有人需要點餐
		w.serveOfCustomerOrder()

		// 移動需要一點時間
		time.Sleep(20 * time.Millisecond)

		// 確認出餐
		w.serveOfDeliverFood()

		// 移動需要一點時間
		time.Sleep(20 * time.Millisecond)
	}
}
