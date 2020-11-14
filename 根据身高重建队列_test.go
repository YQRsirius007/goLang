package pos

import (
	"testing"
)

func TestShouldReturnReceipt_queue(t *testing.T) {
	data := [][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2},
	}
	res := [][]int{
		{5, 0},
		{7, 0},
		{5, 2},
		{6, 1},
		{4, 4},
		{7, 1},
	}
	c := reconstrucQueue(data)

	for i := range c {
		for j, v := range c[i] {
			if v != res[i][j] {
				t.Errorf("want %v, but got %v", res[i], c)
			}
		}

	}
}
