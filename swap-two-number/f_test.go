package swap_two_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var data = []struct {
	a         int
	b         int
	expectedA int
	expectedB int
}{
	{10, 20, 20, 10},
	{2_000, 1_000, 1_000, 2_000},
}

func TestSwap(t *testing.T) {
	for _, data := range data {
		a, b := Swap(data.a, data.b)
		assert.Equal(t, data.expectedA, a)
		assert.Equal(t, data.expectedB, b)
	}
}

func TestSwap2(t *testing.T) {
	for _, data := range data {
		a, b := Swap2(data.a, data.b)
		assert.Equal(t, data.expectedA, a)
		assert.Equal(t, data.expectedB, b)
	}
}

func TestSwap3(t *testing.T) {
	for _, data := range data {
		a, b := Swap3(data.a, data.b)
		assert.Equal(t, data.expectedA, a)
		assert.Equal(t, data.expectedB, b)
	}
}
