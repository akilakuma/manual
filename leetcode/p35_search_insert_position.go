package main

func searchInsert(nums []int, target int) int {
	// binary search ?
	var (
		hitLocal    int
		headerIndex = 0
		eleNum      = len(nums)
	)
	for {
		// 每次範圍砍半
		// 但很麻煩要處理整除的問題，集中右邊處理
		eleNum = eleNum / 2

		if target < nums[headerIndex+eleNum] {

			// min
			if headerIndex+eleNum-1 < 0 {
				hitLocal = 0
				break
			}
			// 往左找
			if nums[headerIndex+eleNum-1] < target {
				hitLocal = headerIndex + eleNum
				break
			}

		} else {
			// 往右找

			// max
			if (headerIndex+eleNum)+1 == len(nums) {
				hitLocal = headerIndex + eleNum + 1
				// 剛好最大也是0的極端case
				if headerIndex+eleNum == 0 && target == nums[0] {
					hitLocal = 0
				}
				// 如果在最右邊，然後跟比較的數相比是等值，要排在前面
				if target == nums[len(nums)-1] {
					hitLocal = len(nums) - 1
				}
				break
			}

			// 已經最左
			if headerIndex+eleNum == 0 && target <= nums[headerIndex+eleNum] {
				hitLocal = headerIndex + eleNum
				break
			}

			// 不左不右剛好卡著
			if nums[headerIndex+eleNum-1] < target && target <= nums[headerIndex+eleNum] {
				hitLocal = headerIndex + eleNum
				break
			}

			// 搜尋點往右縮
			headerIndex += eleNum
			if eleNum%2 == 1 {
				headerIndex += 1
			}
		}
	}
	return hitLocal
}
