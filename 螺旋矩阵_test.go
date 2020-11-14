package pos

import (
	"fmt"
	"testing"
)

func TestShouldReturnReceipt_spiral(t *testing.T) {
	data := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	res := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	c := spiralOrder(data)
	fmt.Println(c, res)

}
