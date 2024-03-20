package swap_two_number

// Swap 使用和交换两个数
func Swap(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

// Swap2 使用亦或交换
func Swap2(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

// Swap3 使用差交换
func Swap3(a, b int) (int, int) {
	a = a - b
	b = a + b
	a = b - a
	return a, b
}
