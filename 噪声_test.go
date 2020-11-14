package pos

import (
	"testing"
)

func TestShouldReturnReceipt_s(t *testing.T) {
	data := [][]int{
		{1, 2, 2, 3, 1},
		{1, 1, 3, 4, 5, 6, 1, 2, 2, 2},
		{1, 2, 2, 3, 1, 4, 2},
		{2, 2, 3, 1, 4, 2},
		{1, 1, 2, 1, 2, 2, 4, 2, 3, 4, 1},
		{1, 1},
		{1, 2},
		{1, 2, 1},
		{1},
	}
	res := []int{
		2,
		3,
		6,
		6,
		6,
		2,
		1,
		3,
		1,
	}
	for i := range data {
		c := findS(data[i])
		if c != res[i] {
			t.Errorf("want %d, but got %d", res[i], c)
		}
	}
}
