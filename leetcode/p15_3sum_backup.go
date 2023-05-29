package main

import (
	"fmt"
	"sort"
)

// Given an integer array nums,
// return all the triplets [nums[i], nums[j],
// nums[k]] such that i != j,
// i != k,
// and j != k,
// and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets
// 用過的位置不能用
func threeSumBackup(nums []int) [][]int {

	// 如果先抓一個元素，就剩兩個元素要處理
	// ex: 第一個元素20，另外2個元素就是-20
	// 那問題會是，怎麼掃過所有元素，抓出兩個加起來是20的組合？

	oriNums := make([]int, len(nums))
	copy(oriNums, nums)

	// 然後塞進map，index=數值，value=位置slice，可能這個數值多次出現
	var numLocateMap = make(map[int][]int)
	for k, v := range oriNums {
		if _, ex := numLocateMap[v]; ex {
			numLocateMap[v] = append(numLocateMap[v], k)
		} else {
			numLocateMap[v] = []int{k}
		}
	}

	// 先sort 由小到大
	sort.Ints(nums)
	//fmt.Println(nums)
	// 標示用slice，slice內容原本都是0
	// 如果已經取用過，則在位置改成1
	// 如果無法配對，改成-1
	var (
		markSlice   = make([]int, len(nums))
		resultSlice [][]int
	)
	for k, v := range nums {
		// 從最左邊開始處理
		var (
			isPairSuccessed bool
		)
		for i := len(nums) - 1; i >= 0; i-- {
			if markSlice[i] == 1 || markSlice[i] == -1 {
				// 被用過或者不能用，則跳過
				continue
			}
			// v：第ㄧ個元素
			// nums[i]：第二個元素
			// lastNum：第三個元素
			lastNum := -(v + nums[i])

			//fmt.Println(v, nums[i], lastNum)

			if _, ex := numLocateMap[lastNum]; ex {
				// 如果相同數字，則需要兩個位置有出現
				c1 := nums[i] == lastNum && len(numLocateMap[lastNum]) < 2
				// 如果3個都是0，處理特殊個案
				c2 := v == 0 && nums[i] == 0 && len(numLocateMap[lastNum]) < 3
				if c1 || c2 {
					continue
				}

				// 存在符合的第三個元素，以及還有尚未使用到的扣打
				isPairSuccessed = true
				// 如果存在的話就湊成三胞胎
				// 被用過的部分標記起來
				markSlice[k] = 1
				markSlice[i] = 1
				// 取第一個來用
				markSlice[numLocateMap[lastNum][0]] = 1

				fmt.Println(v, nums[i], lastNum)

				// 取得原始location
				resultSlice = append(resultSlice, []int{
					numLocateMap[v][0],
					numLocateMap[nums[i]][0],
					numLocateMap[lastNum][0],
				})

				// 已用過就砍掉location資訊
				numLocateMap[v] = numLocateMap[v][1:]
				numLocateMap[nums[i]] = numLocateMap[nums[i]][1:]
				numLocateMap[lastNum] = numLocateMap[lastNum][1:]

				// 如果位置耗盡，連key砍掉
				if len(numLocateMap[v]) == 0 {
					delete(numLocateMap, v)
				}
				if len(numLocateMap[nums[i]]) == 0 {
					delete(numLocateMap, nums[i])
				}
				if len(numLocateMap[lastNum]) == 0 {
					delete(numLocateMap, lastNum)
				}
				break
			}
			if !isPairSuccessed {
				markSlice[k] = -1
			}
		}
	}

	return resultSlice
}

// Given an integer array nums,
// return all the triplets [nums[i], nums[j],
// nums[k]] such that i != j,
// i != k,
// and j != k,
// and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets
// 提供位置的
func threeSumbackup2(nums []int) [][]int {

	// 如果先抓一個元素，就剩兩個元素要處理
	// ex: 第一個元素20，另外2個元素就是-20
	// 那問題會是，怎麼掃過所有元素，抓出兩個加起來是20的組合？

	oriNums := make([]int, len(nums))
	copy(oriNums, nums)

	// 然後塞進map，index=數值，value=位置slice，可能這個數值多次出現
	var numLocateMap = make(map[int][]int)
	for k, v := range oriNums {
		if _, ex := numLocateMap[v]; ex {
			numLocateMap[v] = append(numLocateMap[v], k)
		} else {
			numLocateMap[v] = []int{k}
		}
	}

	// 先sort 由小到大
	sort.Ints(nums)
	//fmt.Println(nums)
	// 標示用slice，slice內容原本都是0
	// 如果無法配對，改成-1
	var (
		markSlice   = make([]int, len(nums))
		resultSlice [][]int
	)
	for k, v := range nums {
		// 從最左邊開始處理
		var (
			isPairSuccessed bool
		)
		for i := len(nums) - 1; i >= 0; i-- {
			if markSlice[i] == -1 {
				// 不能用，則跳過
				continue
			}
			// z1：第ㄧ個元素
			// z2：第二個元素
			// z3：第三個元素
			var (
				z1 = v
				z2 = nums[i]
				z3 = -(v + nums[i])
			)

			//fmt.Println(v, nums[i], lastNum)

			if _, ex := numLocateMap[z3]; ex {

				// 如果3個都是0，處理特殊個案
				c0 := v == 0 && nums[i] == 0 && len(numLocateMap[z3]) < 3
				// 如果相同數字，則需要兩個位置有出現
				c1 := z1 == z2 && len(numLocateMap[z1]) < 2
				c2 := z2 == z3 && len(numLocateMap[z2]) < 2
				c3 := z1 == z3 && len(numLocateMap[z3]) < 2
				// 不符條件，跳出
				if c0 || c1 || c2 || c3 {
					continue
				}

				// 存在符合的第三個元素
				isPairSuccessed = true
				// 如果存在的話就湊成三胞胎
				var (
					local1 = numLocateMap[z1][0]
					local2 = numLocateMap[z2][0]
					local3 = numLocateMap[z3][0]
				)
				if c0 {
					local2 = numLocateMap[z1][1]
					local3 = numLocateMap[z1][2]
				}
				if c1 {
					local2 = numLocateMap[z1][1]
				}
				if c2 {
					local3 = numLocateMap[z2][1]
				}
				if c3 {
					local3 = numLocateMap[z1][1]
				}

				// 取得原始location
				resultSlice = append(resultSlice, []int{
					local1,
					local2,
					local3,
				})

				break
			}
			if !isPairSuccessed {
				markSlice[k] = -1
			}
		}
	}

	return resultSlice
}
