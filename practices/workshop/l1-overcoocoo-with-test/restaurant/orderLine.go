package restaurant

import "sync"

// OrderLine 點餐/出菜桌
type OrderLine struct {
	markRWLocker    *sync.RWMutex         // 讀寫鎖
	markMap         map[string][]FoodMark // 點餐單
	jobs            chan string           // 主菜工作序列
	outputFoodTable chan string           // 出菜檯
}

// FoodMark 餐點註記(點餐單)
type FoodMark struct {
	cID       int32
	food      string
	seatIndex int
}

// NewOrderLine 新的 點餐/出菜桌
func NewOrderLine(chiefs []*Chief) OrderLine {
	o := OrderLine{
		markRWLocker:    new(sync.RWMutex),
		markMap:         make(map[string][]FoodMark, 0),
		jobs:            make(chan string, 100),
		outputFoodTable: make(chan string, 100),
	}

	for _, ch := range chiefs {
		// 讓廚師知道出菜檯位置
		ch.setOutputFoodTable(o.outputFoodTable)
		// 請廚師準備接訂單
		go ch.getOrder(o.jobs)
	}

	return o
}

// income 下單
func (o *OrderLine) income(seat int, cID int32, food string) {

	newOrder := FoodMark{
		cID:       cID,
		food:      food,
		seatIndex: seat,
	}

	// 加到點餐紀錄中
	if _, exists := o.markMap[food]; exists {
		o.markMap[food] = append(o.markMap[food], newOrder)
	} else {
		o.markMap[food] = []FoodMark{
			newOrder,
		}
	}

	// 送工作給廚師
	o.jobs <- food
}
