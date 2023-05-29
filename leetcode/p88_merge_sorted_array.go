package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {

	var (
		spaceLocationIndex = m
		j                  int
	)
	if m == 0 {
		spaceLocationIndex = 0
	}
	if n == 0 {
		return
	}
	// [4 0 0 0 0 0] 1    [1 2 3 5 6] 5
	for i := spaceLocationIndex; i < m+n; i++ {
		fmt.Println(i, j)
		nums1[i] = nums2[j]
		j++
	}
	sort.Ints(nums1)

	fmt.Println(nums1)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	var (
		newArr []int
		n1, n2 int
	)
	for m != 0 && n != 0 {
		fmt.Println("m=", m, n1, "n=", n, n2, "======")
		if nums1[n1] <= nums2[n2] {
			newArr = append(newArr, nums1[n1])
			n1++
			m--
		} else {
			newArr = append(newArr, nums2[n2])
			n2++
			n--
		}
		fmt.Println(newArr)
	}
	fmt.Println("m=", m, n1, "n=", n, n2)
	if m == 0 {
		for i := 0; i < n; i++ {
			fmt.Println("i= ", i)
			newArr = append(newArr, nums2[n2])
			n2++
		}
	}
	if n == 0 {
		for i := 0; i < m; i++ {
			newArr = append(newArr, nums1[n1])
			n1++
		}
	}

	nums1 = nil
	nums1 = make([]int, m+n)
	for i := 0; i < m+n; i++ {
		nums1[i] = newArr[i]
	}
	nums1 = newArr
	fmt.Println(nums1)
}
