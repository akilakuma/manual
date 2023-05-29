package main

import "fmt"

func getSkyline(buildings [][]int) [][]int {
	// 先分group
	// 在分每個group的頭尾
	// 然後再對每個group做處理
	var (
		groupCnt int
		bMaps    = make(map[int]*BuildingStruct)
	)

	// [lefti, righti, heighti]
	for k, v := range buildings {
		// {2,9,10},  {3, 7, 15},
		if k == 0 {
			bMaps[groupCnt] = &BuildingStruct{
				start:        v[0], // 起始點不會變動
				end:          v[1], // 終點可能會延長
				lastHeight:   v[2],
				lastHeader:   v[0], // 如果重疊就必須要重繪
				lastTail:     v[1], // 前一個的屁股，如果下一個建築比較矮，則需要用到
				pairLocation: [][]int{{v[0], v[2]}},
			}
		} else if k < len(buildings) {
			// ## 頭尾都在group的範圍內
			if v[0] < bMaps[groupCnt].end && v[1] < bMaps[groupCnt].end {
				// 如果比前面的高，就要畫天際線
				if v[2] > bMaps[groupCnt].lastHeight {
					bMaps[groupCnt].pairLocation = append(bMaps[groupCnt].pairLocation, []int{v[0], v[2]})
				}
				// ## 頭有在前一個尾的範圍，但是尾巴更往右延伸，但是頭部重疊
			} else if v[0] <= bMaps[groupCnt].end && v[1] > bMaps[groupCnt].end && v[0] != bMaps[groupCnt].lastHeader {
				// 更新尾巴的點
				bMaps[groupCnt].end = v[1]
				// 如果比前面的高，就要畫天際線
				if v[2] > bMaps[groupCnt].lastHeight {
					bMaps[groupCnt].pairLocation = append(bMaps[groupCnt].pairLocation, []int{v[0], v[2]})
				}
				// 如果比前面的矮，也要畫天際線
				if v[2] < bMaps[groupCnt].lastHeight {
					bMaps[groupCnt].pairLocation = append(bMaps[groupCnt].pairLocation, []int{bMaps[groupCnt].lastTail, v[2]})
				}
				// ## 如果頭重疊
			} else if v[0] == bMaps[groupCnt].lastHeader {

				// 天際線必須要重繪
				l := len(bMaps[groupCnt].pairLocation)
				fmt.Println("======", v[2], bMaps[groupCnt].pairLocation[l-1][1], bMaps[groupCnt].pairLocation)
				if v[2] > bMaps[groupCnt].pairLocation[l-1][1] {
					bMaps[groupCnt].pairLocation = append(bMaps[groupCnt].pairLocation[0:l-1], []int{v[0], v[2]})
				}
				// 如果比前面的矮，也要畫天際線
				if v[2] < bMaps[groupCnt].lastHeight && v[1] > bMaps[groupCnt].end {
					bMaps[groupCnt].pairLocation = append(bMaps[groupCnt].pairLocation, []int{bMaps[groupCnt].lastTail, v[2]})
				}

				// 如果尾巴更長
				if v[1] > bMaps[groupCnt].end {
					// 更新尾巴的點
					bMaps[groupCnt].end = v[1]
				}
				// ## 有斷點，變成一個新的group的頭
			} else if v[0] > bMaps[groupCnt].end {
				groupCnt++
				bMaps[groupCnt] = &BuildingStruct{
					start:        v[0], // 起始點不會變動
					end:          v[1], // 終點可能會延長
					lastHeight:   v[2],
					lastHeader:   v[0], // 如果重疊就必須要重繪
					lastTail:     v[1], // 前一個的屁股，如果下一個建築比較矮，則需要用到
					pairLocation: [][]int{{v[0], v[2]}},
				}
			}

			bMaps[groupCnt].lastHeader = v[0]
			bMaps[groupCnt].lastTail = v[1]
			bMaps[groupCnt].lastHeight = v[2]

		}
	}
	r := reBuildGroupData(bMaps)
	var newResult [][]int

	for i := 0; i < len(r); i++ {
		if i < len(r)-1 {
			// 如果兩個一樣就需要合併
			// 比較高，以前面為主
			if r[i][1] == r[i+1][1] {
				newResult = append(newResult, r[i])
				// 跳過下一個
				i++
				continue
			}
			// 比較矮，以後面為主
			if r[i][0] == r[i+1][0] {
				newResult = append(newResult, r[i+1])
				// 跳過下一個
				i++
				continue
			}
		}
		newResult = append(newResult, r[i])
	}
	return newResult
}

func reBuildGroupData(bMaps map[int]*BuildingStruct) [][]int {
	var locateSlice [][]int
	for i := 0; i < len(bMaps); i++ {
		bMaps[i].pairLocation = append(bMaps[i].pairLocation, []int{bMaps[i].end, 0})
		// 每個group 就需要重新整併
		var (
			newGroupSlice [][]int
			heightMax     int
			maxNextLeft   int // 最高之後的下一個點x座標
			maxNextHeight int // 最高之後的下一個點h高度
		)
		for j := 0; j < len(bMaps[i].pairLocation); j++ {
			if heightMax == 0 {
				heightMax = bMaps[i].pairLocation[j][1]
				newGroupSlice = append(newGroupSlice, bMaps[i].pairLocation[j])
				continue
			}
			if bMaps[i].pairLocation[j][1] > heightMax && j < len(bMaps[i].pairLocation)-1 {
				maxNextLeft = bMaps[i].pairLocation[j+1][0]
				maxNextHeight = bMaps[i].pairLocation[j+1][1]
				fmt.Println(maxNextLeft, maxNextHeight)
			}

		}
		locateSlice = append(locateSlice, newGroupSlice...)
	}
	return locateSlice
}

type BuildingStruct struct {
	start        int
	end          int
	lastHeight   int // 前一棟建築物的高度
	lastHeader   int
	lastTail     int
	pairLocation [][]int
}
