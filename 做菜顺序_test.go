package pos

import (
	"testing"
)

func TestShouldReturnReceipt_maxSatisfaction(t *testing.T) {
	data := [][]int{
		{-1, -8, 0, 5, -9},
		{4, 3, 2},
		{-1, -4, -5},
		{-2, 5, -1, 0, 3, -3},
		{76, 83, 55, -36, -8, 40, -60, -72, 27, -32, 37, 1, 77, 24, -46, -26, 20, -89, -35, -99, 58, -7},
	}
	res := []int{
		14,
		20,
		0,
		35,
		7281,
	}
	for i := range data {
		c := maxSatisfaction(data[i])
		if c != res[i] {
			t.Errorf("want %d, but got %d", res[i], c)
		}
	}
}
