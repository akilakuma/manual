package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var firstNode *TreeNode
	firstNode = sliceSplitToNode(nums)

	return firstNode
}

// sliceSplitToNode 給一個slice，然後提供中間當node，以及新分割的左slice和右slice
func sliceSplitToNode(nums []int) *TreeNode {

	var (
		newNode  *TreeNode
		sliceNum = len(nums)
	)
	if sliceNum == 0 {
		return nil
	}
	if sliceNum == 1 {
		newNode = &TreeNode{
			Val: nums[0],
		}
		return newNode
	}
	//fmt.Println("nums:", nums)

	// 是奇數
	var medianIndex = sliceNum / 2
	if sliceNum%2 == 1 {
		medianIndex = sliceNum / 2 // index從0開始，剛好不用+1
	}
	newNode = &TreeNode{
		Val: nums[medianIndex], // 塞中間元素
	}
	//fmt.Println("medianIndex:", medianIndex, "median:", nums[medianIndex])

	// 想像有9個，對切就是0,1,2,3 和 5,6,7,8
	leftSlice := nums[:medianIndex]
	rightSlice := nums[medianIndex+1:]
	//fmt.Println("leftSlice:", leftSlice, " rightSlice:", rightSlice)
	newNode.Left = sliceSplitToNode(leftSlice)
	newNode.Right = sliceSplitToNode(rightSlice)

	return newNode
}
