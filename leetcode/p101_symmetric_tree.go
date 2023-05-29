package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	// 每層存放在一個queue
	// 每一層走完，把content pop 出來一一做比對
	// left tree的內容從左塞到右，right tree的內容從右塞到左

	if root == nil || (root != nil && root.Left == nil && root.Right == nil) {
		return true
	}

	var (
		leftValQueue             []int
		leftParentNodeQueue      []*TreeNode
		nextLeftParentNodeQueue  []*TreeNode
		rightValQueue            []int
		rightParentNodeQueue     []*TreeNode
		nextRightParentNodeQueue []*TreeNode
	)

	// 處理第二層
	if (root.Left == nil && root.Right != nil) || (root.Left != nil && root.Right == nil) {
		// 兩邊要同時有或沒有，不同步return false
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left.Val != root.Right.Val {
		return false
	}

	// 第三層以後開始處理
	leftParentNodeQueue = append(leftParentNodeQueue, root.Left)
	rightParentNodeQueue = append(rightParentNodeQueue, root.Right)

	for len(leftParentNodeQueue) != 0 || len(rightParentNodeQueue) != 0 {

		// 左右數量不對直接再見
		if len(leftParentNodeQueue) != len(rightParentNodeQueue) {
			return false
		}
		fmt.Println(leftParentNodeQueue[0].Val, rightParentNodeQueue[0].Val)

		// 這次遍歷的node都是下一層的parent node
		// 左邊
		for _, leftN := range leftParentNodeQueue {
			fmt.Println("leftN,", leftN.Val)
			// 從左開始
			if leftN.Left != nil {
				leftValQueue = append(leftValQueue, leftN.Left.Val)
				nextLeftParentNodeQueue = append(nextLeftParentNodeQueue, leftN.Left)
			} else {
				leftValQueue = append(leftValQueue, -101)
			}

			if leftN.Right != nil {
				leftValQueue = append(leftValQueue, leftN.Right.Val)
				nextLeftParentNodeQueue = append(nextLeftParentNodeQueue, leftN.Right)
			} else {
				leftValQueue = append(leftValQueue, -101)
			}
		}
		fmt.Println()

		// 右邊
		for _, rightN := range rightParentNodeQueue {
			fmt.Println("rightN,", rightN.Val)
			// 從右開始
			if rightN.Right != nil {
				rightValQueue = append(rightValQueue, rightN.Right.Val)
				nextRightParentNodeQueue = append(nextRightParentNodeQueue, rightN.Right)
			} else {
				rightValQueue = append(rightValQueue, -101)
			}

			if rightN.Left != nil {
				rightValQueue = append(rightValQueue, rightN.Left.Val)
				nextRightParentNodeQueue = append(nextRightParentNodeQueue, rightN.Left)
			} else {
				rightValQueue = append(rightValQueue, -101)
			}
		}

		for i := 0; i < len(leftValQueue); i++ {
			if leftValQueue[i] != rightValQueue[i] {
				return false
			}
		}
		// 檢查完就可以重置了
		leftValQueue = []int{}
		rightValQueue = []int{}
		leftParentNodeQueue = nextLeftParentNodeQueue
		rightParentNodeQueue = nextRightParentNodeQueue
		nextLeftParentNodeQueue = []*TreeNode{}
		nextRightParentNodeQueue = []*TreeNode{}
		//time.Sleep(2 * time.Second)
	}

	return true
}
