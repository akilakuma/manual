package main

import (
	"sort"
	"strconv"
	"strings"
)

// Given an integer array nums,
// return all the triplets [nums[i], nums[j],
// nums[k]] such that i != j,
// i != k,
// and j != k,
// and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets
func threeSum(nums []int) [][]int {

	// 如果先抓一個元素，就剩兩個元素要處理
	// ex: 第一個元素20，另外2個元素就是-20
	// 那問題會是，怎麼掃過所有元素，抓出兩個加起來是20的組合？

	// 然後塞進map，index=數值，value=位置slice，可能這個數值多次出現
	var (
		numLocateMap = make(map[int][]int)
		numCountsMap = make(map[int]int)
		newNums      []int // 其實0超過3個，其他同樣數字超過2個，就可以drop掉
	)
	for k, v := range nums {
		if cnt, ex := numCountsMap[v]; ex && cnt >= 3 {
			continue
		}
		if _, ex := numLocateMap[v]; ex {
			numLocateMap[v] = append(numLocateMap[v], k)
		} else {
			numLocateMap[v] = []int{k}
		}
		numCountsMap[v]++
		newNums = append(newNums, v)
	}

	var (
		markSlice   = make([]int, len(newNums)) // 標示用slice，slice內容原本都是0, 如果無法配對，改成-1
		recordMap   = make(map[string][]int)
		resultSlice [][]int
	)
	for k, v := range newNums {
		// 從最左邊開始處理
		var (
			isPairSuccessed bool
		)
		for i := len(newNums) - 1; i >= 0; i-- {
			if markSlice[i] == -1 {
				// 不能用，則跳過
				continue
			}
			// z1：第ㄧ個元素
			// z2：第二個元素
			// z3：第三個元素
			var (
				z1 = v
				z2 = newNums[i]
				z3 = -(z1 + z2)
			)
			if _, ex := numLocateMap[z3]; ex {

				// 如果3個都是0，處理特殊個案
				c0 := z1 == 0 && z2 == 0 && len(numLocateMap[z3]) < 3
				// 如果相同數字，則需要兩個位置有出現
				c1 := z1 == z2 && len(numLocateMap[z1]) < 2
				c2 := z2 == z3 && len(numLocateMap[z2]) < 2
				c3 := z1 == z3 && len(numLocateMap[z3]) < 2
				// 不符條件，跳出
				if c0 || c1 || c2 || c3 {
					continue
				}
				s := []string{
					strconv.Itoa(z1),
					strconv.Itoa(z2),
					strconv.Itoa(z3),
				}
				sort.Strings(s)
				combineS := strings.Join(s, ",")
				// skip duplicate
				recordMap[combineS] = []int{z1, z2, z3}

				// 存在符合的第三個元素
				isPairSuccessed = true
			}
			if !isPairSuccessed {
				markSlice[k] = -1
			}
		}
	}
	for _, v := range recordMap {
		resultSlice = append(resultSlice, v)
	}

	return resultSlice
}
