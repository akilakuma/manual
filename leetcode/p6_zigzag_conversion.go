package main

// Constraints:
// 1 <= s.length <= 1000
// s consists of English letters (lower-case and upper-case), ',' and '.'.
// 1 <= numRows <= 1000
func convert(s string, numRows int) string {

	// numRows = 一欄最長長度
	// Z字中間一個字的欄數，ex numRows=5, median =5-2=3

	// consider numRows = 1
	// "ABC" numRows = 3
	if numRows == 1 || len(s) == numRows {
		return s
	} else {
		var (
			rowMap   [][]rune
			oneCol   []rune
			colIndex int
		)
		for _, ru := range s {
			var limit int
			if colIndex%(numRows-1) == 0 {
				// 一欄有幾個
				limit = numRows
			} else {
				// 之字形中間都只塞一個
				limit = 1
			}
			if len(oneCol) <= limit {
				oneCol = append(oneCol, ru)
			}
			// 該欄已經塞滿,numRow / 1
			if len(oneCol) == limit {
				rowMap = append(rowMap, oneCol)
				colIndex++
				oneCol = []rune{}
			}
		}
		if len(oneCol) > 0 {
			rowMap = append(rowMap, oneCol)
		}

		return combineZigzag(rowMap, numRows)
	}
}

func combineZigzag(rMap [][]rune, numRows int) string {
	var (
		transMap = make([][]rune, numRows)
		finalRu  []rune
	)
	//fmt.Println("len of rMap =", len(rMap))

	for key, value := range rMap {
		// 完整欄，全塞
		if key == 0 || key%(numRows-1) == 0 {
			for i := 0; i < numRows; i++ {

				if len(value) > i {
					//fmt.Println(key, i, string(value[i]))
					transMap[i] = append(transMap[i], value[i])
				}
			}
		} else {
			// 非完整欄，找對位置塞
			// 這種slice只會有一個字
			reciprocal := (numRows - 1) - (key % (numRows - 1))
			//fmt.Println(key, reciprocal)
			transMap[reciprocal] = append(transMap[reciprocal], value[0])
		}
	}

	for _, ru := range transMap {
		finalRu = append(finalRu, ru...)
	}

	return string(finalRu)
}
