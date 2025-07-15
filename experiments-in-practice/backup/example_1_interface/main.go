package main

// AppleFamily 符合iphone定義
type AppleFamily interface {
	showAppleMark()
	hasSiri()
	priceOverTwentyThousands()
}

// fashionAppleProduct 超酷的蘋果產品
func fashionAppleProduct(a AppleFamily) {
	a.showAppleMark()
	a.hasSiri()
	a.priceOverTwentyThousands()
}

type job interface {
	Run()
}

func AddJob(j job) {
	j.Run()
}

type NewJob struct {
	pid         int
	overlapping bool
	isRunning   bool
}

func (n *NewJob) Run() {
	if n.overlapping {
		// 多一個執行
	}

	if n.isRunning {
		return
	}

	n.isRunning = true

	executeCalculation()

	n.isRunning = false

}

type OriginJob struct {
}

func (o *OriginJob) Run() {
	executeCalculation()
}

// executeCalculation 執行計算
func executeCalculation() {
	//
}

func main() {

}
