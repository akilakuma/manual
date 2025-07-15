package restaurant

import (
	"errors"
	"log"
	"sync"
	"sync/atomic"
)

// Restaurant 餐廳
type Restaurant struct {
	customerMaxID int32                  // 客人編號(依據該編號依序送餐)
	maxSeats      int                    // 最大座位數
	currentSeats  int                    // 目前使用的座位數
	seatRWlocker  *sync.RWMutex          // 座位的讀寫鎖
	chiefs        map[*Chief]bool        // 廚師們
	waiter        map[*Waiter]bool       // 服務生們
	receptionist  map[*Receptionist]bool // 接待人員
	orderLine     OrderLine              // 點餐/出菜桌
	seats         map[int]Seat           // 座位，key:座位編號，value:座位
}

// Seat 座位資訊
type Seat struct {
	cID      int32     // 客人編號
	customer *Customer // 客人
	food     string    // 餐點
}

// CreateRestaurant 實例化一個新餐廳
func CreateRestaurant() *Restaurant {
	var maxSeats = 10
	r := &Restaurant{
		maxSeats:     maxSeats,
		seatRWlocker: new(sync.RWMutex),
		chiefs:       make(map[*Chief]bool, 0),
		waiter:       make(map[*Waiter]bool, 0),
		receptionist: make(map[*Receptionist]bool, 0),

		seats: make(map[int]Seat, maxSeats),
	}

	go r.workerOnDuty()

	return r
}

// NewCustomerID 的新客人識別ID
func (r *Restaurant) NewCustomerID() int32 {
	return atomic.AddInt32(&r.customerMaxID, 10)
}

// workerOnDuty 餐廳人員上班
func (r *Restaurant) workerOnDuty() {

	var chiefSlice []*Chief
	for i := 0; i < 2; i++ {
		c := NewChief()
		r.chiefs[c] = true
		chiefSlice = append(chiefSlice, c)
	}

	for i := 0; i < 2; i++ {
		w := NewWaiter(r)
		r.waiter[w] = true
	}

	for i := 0; i < 1; i++ {
		receptor := NewReceptionist(r)
		r.receptionist[receptor] = true
	}

	// 點餐/出餐檯 置入上工的廚師
	r.orderLine = NewOrderLine(chiefSlice)

	log.Println("餐廳工作人員上工就緒")
	log.Println()
}

// CustomerComein 客人進入
func (r *Restaurant) CustomerComein(c *Customer) error {

	var freeReceptor *Receptionist
	// 找出目前有空的接待人員
	for receptor := range r.receptionist {

		if !receptor.isServing {
			freeReceptor = receptor
		}
	}
	if freeReceptor == nil {
		return errors.New("目前沒有空的接待人員")
	}

	// 接待人員帶位
	err := freeReceptor.leadToSeat(c)
	if err != nil {
		return err
	}

	return nil
}
