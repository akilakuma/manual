package main

func sortColors(nums []int) {

	var (
		lastIndex int
		lastV     int
	)
	for i := 0; i < len(nums); i++ {
		//fmt.Println(i, nums)
		if i == 0 {
			lastIndex = 0
			lastV = nums[i]
			continue
		}
		// 如果比前面的小，就交換
		if nums[i] < lastV {
			tmp := nums[i]
			nums[i] = lastV
			nums[lastIndex] = tmp
			i = -1
		} else {
			lastIndex = i
			lastV = nums[i]
		}
	}
}
