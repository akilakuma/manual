package main

func moveZeroes(nums []int) {

	var (
		OnRightSide  bool
		nonZeroIndex int
		nonZeroCnt   int
	)
	for i := len(nums) - 1; i >= 0; i-- {

		// 從屁股數來，如果
		// 1.是0，但右邊沒非0的數就繼續，就繼續往左數
		// 2.是0，但右邊有非0的數，開始區段交換
		if nums[i] == 0 && !OnRightSide {
			continue
		} else if nums[i] == 0 && OnRightSide {

			for j := i; j < nonZeroIndex; j++ {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
			}
			// 非零數一起往左移動一格
			nonZeroIndex = nonZeroIndex - 1

		} else {
			nonZeroCnt++
			// 記得從哪裡

			OnRightSide = true
			if nonZeroIndex == 0 {
				nonZeroIndex = i
			}
		}
	}
}
