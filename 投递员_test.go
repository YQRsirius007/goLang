package pos

import (
	"testing"
)

func TestShouldReturnReceipt_post(t *testing.T) {
	data := [][]int{
		{1, 0, 1},
		{1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1},
	}
	res := []int{
		1,
		3,
		4,
	}
	for i := range data {
		c := postman(data[i])
		if c != res[i] {
			t.Errorf("want %d, but got %d", res[i], c)
		}
	}
}
