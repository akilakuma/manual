package main

func maxArea(height []int) int {

	var (
		max int
	)

	for k, v := range height {

		// 從最遠數過來
		for i := len(height) - 1; i >= 1; i-- {
			// 間隔數
			distance := i - k
			// v *距離決定了最大面積多少，所以可以知道接下來往左移動一格，最大面積
			nextThresholdArea := (distance - 1) * v

			var area int
			// 比v還高，以v計算
			if height[i] >= v {
				area = v * distance
			} else {
				area = height[i] * distance
			}

			if area > max {
				max = area
			}
			// 現在的面積比下個最大可能面積還高，就不用看了
			if area > nextThresholdArea {
				break
			}
			if nextThresholdArea < max {
				break
			}
		}
	}
	return max
}
