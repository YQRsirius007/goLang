package pos

import (
	"testing"
)

func TestShouldReturnReceipt_jump1(t *testing.T) {
	data := [][]int{
		{2, 3, 1, 1, 4},
		{3, 2, 1, 0, 4},
	}
	dataN := []int{
		5,
		5,
	}
	res := []bool{
		true,
		false,
	}
	for i := range data {
		c := jump1(data[i], dataN[i])
		if c != res[i] {
			t.Errorf("want %v, but got %v", res[i], c)
		}
	}
}

func TestShouldReturnReceipt_jump(t *testing.T) {
	data := make([][]int, 1)
	data[0] = []int{2, 3, 1, 1, 4}
	res := []int{
		2,
	}
	for i := range data {
		c := jump(data[i])
		if c != res[i] {
			t.Errorf("want %v, but got %v", res[i], c)
		}
	}
}
