package main

import "fmt"

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	var (
		newHead       *ListNode
		leftLastNode  *ListNode
		rightLastNode *ListNode
		leftSlice     []*ListNode
		rightSlice    []*ListNode
	)

	for {
		if head.Val < x {
			leftSlice = append(leftSlice, head)
		} else {
			rightSlice = append(rightSlice, head)
		}

		if head.Next != nil {
			head = head.Next
		} else {
			break
		}
	}

	fmt.Println(len(leftSlice), len(rightSlice))
	if len(leftSlice) > 0 {

		newHead = leftSlice[0]
		leftLastNode = newHead
		leftLastNode.Next = nil

		for i := 0; i < len(leftSlice)-1; i++ {
			leftSlice[i].Next = leftSlice[i+1]
			leftLastNode = leftSlice[i+1]
			leftLastNode.Next = nil
		}

		if len(rightSlice) > 0 {
			if leftLastNode != nil {
				leftLastNode.Next = rightSlice[0]
				rightSlice[0].Next = nil
			}
		}

	} else {
		// right 應該一定有東西
		newHead = rightSlice[0]
		rightLastNode = newHead
		rightLastNode.Next = nil
	}

	for i := 0; i < len(rightSlice)-1; i++ {
		rightSlice[i].Next = rightSlice[i+1]
		rightLastNode = rightSlice[i+1]
		rightLastNode.Next = nil
	}
	p := newHead
	for {
		fmt.Println(p.Val)

		if p.Next != nil {
			p = p.Next
		} else {
			break
		}
	}

	return newHead
}
