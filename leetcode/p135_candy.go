package main

import "fmt"

func candy(ratings []int) int {
	/*
		// 三個一組
		// 由右往左推
		// 往左推幾個沒極限，一路推到底
		上面是rating
		下面是給的糖果數

		1 2
		-------------
		1 2 3
		1 2 3
		=====
		1 2 2
		1 2 1
		=====
		1 2 1
		1 2 1
		=====


		1 1
		-------------
		1 1 2
		1 1 2
		=====
		1 1 1
		1 1 1
		=====
		1 1 0
		1 2 1
		=====

		2 1
		-------------
		2 1 2
		2 1 2
		=====
		2 1 1
		2 1 1
		=====
		2 1 0
		3 2 1
		=====
	*/

	var candySlice = make([]int, len(ratings))

	if len(ratings) == 1 {
		return 1
	}

	// 只要跟別人比較
	for k, v := range ratings {
		// 以自己為最右邊，只往左看

		// 第一個塞1即可
		if k == 0 {
			candySlice[0] = 1
			continue
		}
		fmt.Println("k=", k, "v=", v, candySlice)
		// 從第2個之後就要做比較
		if v > ratings[k-1] {
			candySlice[k] = candySlice[k-1] + 1 // 比左邊厲害，+1
		} else if v == ratings[k-1] {
			candySlice[k] = 1
		} else {
			// 小於左邊，最麻煩
			candySlice[k] = 1
			if k >= 2 && ratings[k-2] >= ratings[k-1] {
				candySlice[k-1] += 1
			} else if k <= 1 {
				candySlice[k-1] += 1
			}
			for i := k - 1; i > 0; i-- {
				if ratings[i] >= ratings[i-1] {
					// 什麼事都不用做，而且不需要往左推算
					break
				} else {
					if i >= 2 && ratings[i-2] >= ratings[i-1] {
						candySlice[i-1] += 1
					} else if i <= 1 {
						candySlice[i-1] += 1
					} else if candySlice[i-1] <= candySlice[i] {
						if ratings[i-1] > ratings[i] {
							candySlice[i-1] = candySlice[i] + 1
						}
					}
				}
			}
		}
	}

	var total int
	for _, v := range candySlice {
		fmt.Print(v, "")
		total += v
	}
	fmt.Println()
	return total
}
