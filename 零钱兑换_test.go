package pos

import (
	"testing"
)

func TestShouldReturnReceipt_coin(t *testing.T) {
	data := [][]int{
		{1, 2, 5},
		{2},
	}
	dataN := []int{
		11,
		3,
	}
	res := []int{
		3,
		-1,
	}
	for i := range data {
		c := coinChange(data[i], dataN[i])
		if c != res[i] {
			t.Errorf("want %v, but got %v", res[i], c)
		}
	}
}
