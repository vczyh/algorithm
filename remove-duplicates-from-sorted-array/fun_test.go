package remove_duplicates_from_sorted_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n := removeDuplicates2(nums)
	t.Log(n)
}

// 使用 BitSet ，数据量小的时候，内存占用太高
func removeDuplicates(nums []int) int {
	// bit set
	bsLen := (2*10_000+1-1)/64 + 1
	words := make([]uint64, bsLen)

	size := len(nums)
	i := 0
	k := 0
	for i < size {
		v := nums[i]
		bv := v + 10_000
		if words[bv/64]&(1<<(bv%64)) != 0 {
			i++
		} else {
			words[bv/64] |= 1 << (bv % 64)
			nums[k] = v
			k++
			i++
		}
	}

	return k
}

func removeDuplicates2(nums []int) int {
	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[k] {
			continue
		}
		k++
		nums[k] = nums[i]
	}
	return k + 1
}

func Test2(t *testing.T) {
	bsLen := (2*10_000+1-1)/64 + 1
	words := make([]uint64, bsLen)
	bv := 10000
	words[bv/64] |= 1 << (bv % 64)
	fmt.Println(bv / 64)
	fmt.Println(bv % 64)

	fmt.Println(words[bv/64])
	fmt.Println(1 << (bv % 64))

	//a := words[bv/64]
	//b := 1 << (bv % 64)
	//fmt.Printf("%b\n", a)
	//fmt.Printf("%b\n", b)
	fmt.Println(words[bv/64] & (1 << (bv % 64)))
}
