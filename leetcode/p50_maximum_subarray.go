package main

import (
	"container/list"
	"fmt"
	"time"
)

/*

-2 lmax = -2 rmax = -2
-2 1 lmax = -1 rmax = 1 (pop -2)
1 -3 lmax = 1 rmax = -2
1 -3 4 lmax = 2 rmax = 4 (pop 1 -3)
4 -1 lmax = 4 rmax = 3
4 -1 2 lmax = 5 rmax = 5
4 -1 2 1 lmax = 6 rmax = 6
4 -1 2 1 -5 lmax = 6 rmax = 1
4 -1 2 1 -5 4 lmax = 6 rmax = 5

lmax 是從左數來，rmax 是從右數來
一旦 rmax > lmax 從左邊開始pop 直到 lmax = rmax

*/

func maxSubArray(nums []int) int {
	fmt.Println("nums = ", len(nums))
	a := time.Now()
	var (
		lmax             int = nums[0]
		rmax             int = nums[0]
		subArrayList         = list.New()
		lboundle         int
		rboundle         int
		isLCollecting    bool
		isRCollecting    bool
		lCnt, rCnt, cCnt int
	)

	var fromLeft = func() {

		if subArrayList.Len() < 1 {
			return
		}

		var (
			lTotal   int
			leftHead = subArrayList.Front()
		)
		for {
			lTotal += leftHead.Value.(int)

			if lTotal > lmax {
				lmax = lTotal
			}
			if leftHead.Next() == nil {
				break
			} else {
				leftHead = leftHead.Next()
			}
		}
	}

	var fromRight = func() {

		if subArrayList.Len() < 1 {
			return
		}

		var (
			rTotal    int
			rightHead = subArrayList.Back()
		)
		for {
			rTotal += rightHead.Value.(int)
			if rTotal > rmax {
				rmax = rTotal
			}
			if rightHead.Prev() == nil {
				break
			} else {
				rightHead = rightHead.Prev()
			}
		}
	}

	for k, v := range nums {

		subArrayList.PushBack(v)

		// 第一筆不管它
		if k == 0 {
			continue
		}

		if lmax > 0 {
			isLCollecting = true
		}

		// 想要後面的正數，必須吃下前面的負數
		if isLCollecting {
			lboundle += v
		}

		if rmax > 0 {
			isRCollecting = true
		}

		// 想要後面的正數，必須吃下前面的負數
		if isRCollecting {
			rboundle += v
		}

		// 處理從左數過來，更新lmax
		if lmax < 0 || (v > 0 && lboundle > 0) {
			if v > 0 && lboundle > 0 {
				lmax += lboundle
			} else {
				fromLeft()
			}

			// 同捆包校正
			lboundle = 0
			lCnt++
		}

		// 處理從右數過來，更新rmax
		// v 要大於lmax和rmax之間的差距，rmax才會比之前大
		if rmax <= 0 || (v > 0 && rboundle > (lmax-rmax)) || v > rmax {
			fromRight()

			// 同捆包校正
			rboundle = 0
			rCnt++
		}

		// 如果lmax 小於rmax，左邊的東西一直丟掉，直到lmax = rmax
		if lmax < rmax {

			for lmax < rmax {
				e := subArrayList.Front()
				subArrayList.Remove(e)

				// 處理從左數過來，更新lmax
				fromLeft()
			}
			// 同捆包校正
			lboundle = 0
			rboundle = 0
			cCnt++
		}
	}
	fmt.Println(time.Now().Sub(a), lCnt, rCnt, cCnt)

	// 如果原本rmax比較大，最後一步也會讓lmax = rmax
	return lmax
}
