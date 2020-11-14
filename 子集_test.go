package pos

import (
	"testing"
)

func TestShouldReturnReceipt_subset(t *testing.T) {
	data := make([][]int, 1)
	data[0] = []int{1, 2, 3}
	res := make([][]int, 1)
	res[0] = []int{}
	for i := range data {
		c := subsets(data[i])
		t.Errorf("want %v, but got %v", res[i], c)
	}
}
func TestShouldReturnReceipt_subset2(t *testing.T) {
	data := make([][]int, 1)
	data[0] = []int{1, 2, 3}
	res := make([][]int, 1)
	res[0] = []int{}
	for i := range data {
		c := subsets(data[i])
		t.Errorf("2want %v, but got %v", res[i], c)
	}
}
