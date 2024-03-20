package main

import (
	"fmt"
)

// https://leetcode.cn/problems/merge-sorted-array/description/

func main() {
	num1 := []int{4, 5, 6, 0, 0, 0}
	num2 := []int{1, 2, 3}
	merge(num1, 3, num2, 3)
	fmt.Println(num1)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	if n == 0 {
		return
	}

	i := m - 1
	j := n - 1
	k := m + n - 1

	for j >= 0 && i >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	fmt.Println(nums1)

	for j >= 0 {
		nums1[j] = nums2[j]
		j--
	}
}
