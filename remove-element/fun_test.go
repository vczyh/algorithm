package remove_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.cn/problems/remove-element/

func TestRemoveElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	size := removeElement(nums, 2)
	assert.Equal(t, 5, size)
}

func removeElement(nums []int, val int) int {
	i, k := 0, 0
	for i < len(nums) {
		if nums[i] == val {
			i++
		} else {
			nums[k] = nums[i]
			i++
			k++
		}
	}
	return k
}
