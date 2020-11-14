package pos

import (
	"fmt"
	"testing"
)

func TestShouldReturnReceipt_combination(t *testing.T) {
	data := [][]int{
		{2, 3, 6, 7},
		{2, 3, 5},
	}
	target := []int{
		7,
		8,
	}
	// res := [][]int{
	// 	{7},
	// 	{2,2,3},
	// }
	for i := range data {
		c := combinationSum(data[i], target[i])
		fmt.Println(c)
		// if c!=res[i] {
		// 	t.Errorf("want %d, but got %d",res[i], c)
		// }
	}
}
