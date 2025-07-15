package main

import (
	"log"

	"github.com/shopspring/decimal"
)

func main() {
	// test1()
	// test3()
	getBet()
}

// func (d Decimal) FloorWithFloat(floor int32) Decimal {

// 	d.exp = d.exp + floor

// 	d.ensureInitialized()

// 	if d.exp >= 0 {
// 		return d
// 	}

// 	exp := big.NewInt(10)
// 	// NOTE(vadim): must negate after casting to prevent int32 overflow
// 	exp.Exp(exp, big.NewInt(-int64(d.exp)), nil)
// 	z := new(big.Int).Div(d.value, exp)

// 	return Decimal{value: z, exp: -1 * floor}
// }

// func (d Decimal) FloorWithFloat(floor int32) Decimal {

// 	d.exp = d.exp + floor

// 	d.ensureInitialized()

// 	if d.exp >= 0 {
// 		return d
// 	}

// 	exp := big.NewInt(10)
// 	// NOTE(vadim): must negate after casting to prevent int32 overflow
// 	exp.Exp(exp, big.NewInt(-int64(d.exp)), nil)
// 	z := new(big.Int).Div(d.value, exp)

// 	return Decimal{value: z, exp: -1 * floor}
// }

func test1() {

	floor := int32(3)
	a := decimal.NewFromFloat(-2.05)

	b := a.Shift(floor).Abs().Floor().Shift(-1 * floor)

	if a.IsNegative() {
		b = b.Neg()
	}

	c, _ := b.Float64()
	log.Println(c)

	// b := a.Mul(decimal.NewFromFloat(100))
	// log.Println(b)

	// if b.IsNegative() {
	// 	tmpNegative := b.Neg()
	// 	c := tmpNegative.Floor()
	// 	finalNegative := c.Neg()
	// 	d := finalNegative.Div(decimal.NewFromFloat(100))
	// 	log.Println(d)
	// } else {
	// 	floor := 2
	// 	c := b.Shift(floor).Floor().Shift(-1 * floor)

	// 	b.Shift(floor).Neg().Floor().Neg().Shift(-1 * floor)

	// 	log.Println(c)
	// 	d := c.Div(decimal.NewFromFloat(100))

	// 	log.Println(d)
	// }

}

// func test2() {
// 	d := -1129.623456

// 	dec1 := decimal.NewFromFloat(d)

// 	result := dec1.FloorWithFloat(2)
// 	fmt.Println(result)
// }

func test3() {

	percentInDetailSettingDecimal := decimal.NewFromFloat(61.1)
	gameBetBalanceDecimal := decimal.NewFromFloat(150.0000)

	tmp, _ := percentInDetailSettingDecimal.Mul(gameBetBalanceDecimal).Div(decimal.NewFromFloat(100)).Float64()
	amount := FloorFloatNum(tmp, 2)

	// 這次可領的錢 = 計算出來的錢 - 之前領的錢
	amount = amount - 0
	amount = FloorFloatNum(amount, 2)

	lastMax := 1989.86
	//判斷返水上限門檻
	if amount >= float64(lastMax) {
		// 上限也到頭了
		amount = float64(lastMax)
		lastMax = 0
	} else {
		log.Println(lastMax)
		// 本次剩餘上限 - 這次可領的錢
		// 若是負的不會因此增加上限
		// (因為你不知道加回的額度會不會大於本次的上限或每日上限，反正這次沒領到，下次可領)
		if amount > 0 {
			amountDecimal := decimal.NewFromFloat(amount)
			lastMaxDecimal := decimal.NewFromFloat(lastMax)
			s, _ := lastMaxDecimal.Sub(amountDecimal).Float64()

			lastMax = FloorFloatNum(s, 2)
		}
	}

	log.Println(lastMax)
}

func FloorFloatNum(f float64, para int) float64 {

	// 將指定的位數轉成int32
	digits := int32(para)
	sourceNum := decimal.NewFromFloat(f)

	// 位移X個位數(10進位)，取絕對值(因為可能是負數)，捨去整數後的數字，再位移回來
	resultNum := sourceNum.Shift(digits).Abs().Floor().Shift(-1 * digits)

	// 如果是負數，需要轉回來(取絕對值的時候變成了正數)
	if sourceNum.IsNegative() {
		resultNum = resultNum.Neg()
	}

	// 不需要管精準度
	resultNumToFloat64, _ := resultNum.Float64()

	return resultNumToFloat64
}


func getBet() {
	gameData := make(map[int]string, 0)

	gameData[0] = "card"
	gameData[1] = "sport"
	gameData[2] = "lottery"
	gameData[3] = "live"
	gameData[4] = "machine"
	gameData[5] = "fish"

	// for index, value := range gameData {
		// 處理資料
	// }
}
