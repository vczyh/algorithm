package sort

// 所有排序算法

// BubbleSort 冒泡排序
// 时间复杂度：O(n^2)
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
}

// SelectSort 选择排序
// 时间复杂度：O(n^2)
func SelectSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[min]
		arr[min] = tmp
	}
}

// InsertSort 插入排序
// 时间复杂度：O(n^2)
func InsertSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		current := arr[i]
		preIndex := i - 1
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
	}
}

// ShellSort 希尔排序 交换元素
// 时间复杂度：O(n^3/2)
func ShellSort(arr []int) {
	n := len(arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			j := i
			pre := j - gap
			for pre >= 0 && arr[pre] > arr[j] {
				// 交换
				arr[pre], arr[j] = arr[j], arr[pre]
				j -= gap
				pre = j - gap
			}
		}
	}
}

// ShellSort2 希尔排序 移动元素
// 时间复杂度：O(n^3/2)
func ShellSort2(arr []int) {
	n := len(arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			j := i
			current := arr[j]
			for j-gap >= 0 && arr[j-gap] > current {
				// 移动
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = current
		}
	}
}

//func ShellSort(arr []int) {
//	n := len(arr)
//	// 初始增量设置为数组长度的一半
//	for gap := n / 2; gap > 0; gap /= 2 {
//		// 从增量位置开始，逐步对每个子序列进行插入排序
//		for i := gap; i < n; i++ {
//			// 保存当前位置的元素，以及预备插入的位置
//			temp := arr[i]
//			j := i
//			// 插入排序的核心逻辑，比较并交换
//			for ; j >= gap && arr[j-gap] > temp; j -= gap {
//				arr[j] = arr[j-gap]
//			}
//			// 将元素插入到预定位置
//			arr[j] = temp
//		}
//	}
//}

// MergeSort 归并排序
func MergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	mid := n / 2
	// 左边归并排序
	left := MergeSort(arr[:mid])
	// 右边归并排序
	right := MergeSort(arr[mid:])
	// 合并
	return mergeSortMerge(left, right)
}

func mergeSortMerge(left, right []int) []int {
	ll := len(left)
	rl := len(right)

	var arr []int
	l, r := 0, 0

	for l < ll && r < rl {
		if left[l] < right[r] {
			arr = append(arr, left[l])
			l++
		} else {
			arr = append(arr, right[r])
			r++
		}
	}

	arr = append(arr, left[l:]...)
	arr = append(arr, right[r:]...)
	return arr
}
