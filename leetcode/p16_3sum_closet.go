package main

import "sort"

// Given an integer array nums,
// return all the triplets [nums[i], nums[j],
// nums[k]] such that i != j,
// i != k,
// and j != k,
// and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets
func threeSumClosest(nums []int, target int) int {
	//t := time.Now()

	sort.Ints(nums)

	var (
		leastSmall = nums[0] + nums[1] + nums[2]
		leastBig   = nums[len(nums)-3] + nums[len(nums)-2] + nums[len(nums)-1]
	)

	for i := 0; i < len(nums); i++ {

		// 如果第二個數都比最大的大，那第一層整個不用看了
		if i+1 < len(nums) && nums[i+1] > 0 && nums[i]+nums[i+1] > leastBig {
			break
		}
		for j := i + 1; j < len(nums); j++ {

			// 如果第三個數都比最大的大，那第二層整個不用看了
			if j+1 < len(nums) && nums[i]+nums[j]+nums[j+1] > leastBig {
				break
			}
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				// 比目標小
				// -4   < -2       -5 < -4
				if sum <= target && leastSmall <= sum {
					leastSmall = sum
				}
				// 比目標大
				// 4 > 2                5 >  4
				if sum >= target && leastBig >= sum {
					leastBig = sum
					// 因為有先排序過，後面只會更大
					break
				}
			}
		}
	}

	//fmt.Println(time.Now().Sub(t))
	if target-leastSmall < leastBig-target {
		return leastSmall
	}
	return leastBig
}

// Given an integer array nums,
// return all the triplets [nums[i], nums[j],
// nums[k]] such that i != j,
// i != k,
// and j != k,
// and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets
func threeSumClosestV1(nums []int, target int) int {
	//t := time.Now()

	// 先暴力n^3看看
	var (
		leastBig   int = target + 20001
		leastSmall int = target - 20001
	)
	if nums[0]+nums[1]+nums[2] > target {
		leastBig = nums[0] + nums[1] + nums[2]
	} else {
		leastSmall = nums[0] + nums[1] + nums[2]
	}

	for xIndex, x := range nums {
		for yIndex, y := range nums {
			for zIndex, z := range nums {
				if xIndex != yIndex && yIndex != zIndex && xIndex != zIndex {
					sum := x + y + z
					// 比目標小
					// -4   < -2       -5 < -4
					if sum <= target && leastSmall <= sum {
						leastSmall = sum
					}
					// 比目標大
					// 4 > 2                5 >  4
					if sum >= target && leastBig >= sum {
						leastBig = sum
					}
				}
			}
		}
	}
	if target-leastSmall < leastBig-target {
		return leastSmall
	}
	return leastBig
}
