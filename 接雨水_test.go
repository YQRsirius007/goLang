package pos

import "testing"

func TestShouldReturnReceipt_trap(t *testing.T) {
	data := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
	}

	res := []int{
		6,
	}
	for i, v := range data {
		c := trap(v)
		if c != res[i] {
			t.Errorf("want %d, but got %d", res[i], c)
		}
	}
}
