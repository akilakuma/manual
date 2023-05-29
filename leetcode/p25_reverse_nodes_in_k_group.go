package main

import (
	"fmt"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 理解題目：
	// 給一串list node
	// 從最左邊開始，每經過指定k個node，對這些node做reverse
	// 最後面如果沒有被k整除，則保持原樣不做reverse

	// memo :
	// reverse好像也不等於排序，他沒說給的list是sort過

	if k == 1 {
		return head
	}

	var (
		kCountDown = k
		// 在每一段的頭，reverse之後會變成最後一個，要記錄來接下一段reverse
		partHeader *ListNode
		// 原本的第一個，會變成reverse之後的tail，接下一個partHeadr
		firstPartTails  *ListNode
		secondPartTails *ListNode
		// 先記錄下來，下一個才知道接誰
		lastNode *ListNode
		// 要找出新的頭
		isFirstRevere bool = true
		// 新的頭
		newHeadNode         *ListNode
		originPartNodeSlice []*ListNode

		eleCnt = 1
	)

	for {
		// 先記錄下來，原本的下一個node是什麼
		// 不然接下來.Next要去連接上一個
		var oriNextNode *ListNode
		if head != nil && head.Next != nil {
			oriNextNode = head.Next
		}

		// 重置
		if kCountDown == 0 {
			kCountDown = k
		}

		if kCountDown == k {
			originPartNodeSlice = []*ListNode{}
			originPartNodeSlice = append(originPartNodeSlice, head)

			// reverse之後的尾巴
			firstPartTails = secondPartTails
			secondPartTails = head
			//}
			lastNode = head
			// 尾巴的nex淨空
			// 如果有新的頭出現自然去接新的頭
			secondPartTails.Next = nil

		} else if 1 < kCountDown && kCountDown < k {

			originPartNodeSlice = append(originPartNodeSlice, head)

			head.Next = lastNode
			lastNode = head

		} else {
			head.Next = lastNode
			partHeader = head

			if isFirstRevere {
				isFirstRevere = false
				newHeadNode = head
				// 第一次則沒有接尾巴
				// 因為要等到第二個頭出現，才知道要怎麼接
			} else {
				// 知道頭在這裡，之前的尾巴來接到這個頭
				firstPartTails.Next = partHeader
			}
		}

		kCountDown--

		if oriNextNode != nil {
			head = oriNextNode

			// 數有幾個node
			eleCnt++
		} else {
			if eleCnt%k > 0 {
				// 遇到無法整除
				var (
					last *ListNode
				)
				for index, v := range originPartNodeSlice {
					if index == 0 {
						firstPartTails.Next = v
						last = v

						continue
					}
					last.Next = v
					last = v

				}
				last.Next = nil
			}
			break
		}
	}

	for newHeadNode != nil {
		fmt.Println(newHeadNode.Val)
		newHeadNode = newHeadNode.Next
	}
	return newHeadNode
}
