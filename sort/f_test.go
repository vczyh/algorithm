package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var data = []struct {
	arr      []int
	expected []int
}{
	{[]int{3, 2, 1, 4, 5}, []int{1, 2, 3, 4, 5}},
	{[]int{3, 2, 1, 2, 4, 5}, []int{1, 2, 2, 3, 4, 5}},
}

func TestBubbleSort(t *testing.T) {
	for _, data := range data {
		BubbleSort(data.arr)
		assert.Equal(t, data.expected, data.arr)
	}
}

func TestSelectSort(t *testing.T) {
	for _, data := range data {
		SelectSort(data.arr)
		assert.Equal(t, data.expected, data.arr)
	}
}

func TestInsertSort(t *testing.T) {
	for _, data := range data {
		InsertSort(data.arr)
		assert.Equal(t, data.expected, data.arr)
	}
}

func TestShellSort(t *testing.T) {
	for _, data := range data {
		ShellSort(data.arr)
		assert.Equal(t, data.expected, data.arr)
	}
}

func TestShellSort2(t *testing.T) {
	for _, data := range data {
		ShellSort2(data.arr)
		assert.Equal(t, data.expected, data.arr)
	}
}
