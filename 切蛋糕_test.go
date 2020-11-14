package pos

import "testing"

func TestShouldReturnReceipt_maxCake(t *testing.T) {
	data1 := []int{
		5,
		5,
		5,
	}
	data2 := []int{
		4,
		4,
		4,
	}
	arr1 := [][]int{
		{1, 2, 4},
		{3, 1},
		{3},
	}
	arr2 := [][]int{
		{1, 3},
		{1},
		{3},
	}
	res := []int{
		4,
		6,
		9,
	}
	for i := range data1 {
		c := maxCake(data1[i], data2[i], arr1[i], arr2[i])
		if c != res[i] {
			t.Errorf("want %d, but got %d", res[i], c)
		}
	}
}
