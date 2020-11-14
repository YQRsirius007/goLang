package pos

import (
	"fmt"
	"testing"
)

// [[8,0], [4,4], [8,1], [5,0], [6,1], [5,2]]
// [[5,0], [8,0], [5,2], [6,1], [4,4], [8,1]]
func TestShouldReturnReceipt_timu2(t *testing.T) {
	s := "hello你好"
	fmt.Println(len(s), len([]rune(s)))

	data1 := [][]int{
		{1, 3},
		{-2, 2},
	}
	res1 := [][]int{
		{-2, 2},
	}
	data2 := [][]int{
		{3, 3}, {5, -1}, {-2, 4}, {10000, 10000},
	}
	fmt.Println(res1)
	fmt.Println(numClosestPositions(data1, 2))
	fmt.Println(numClosestPositions(data2, 2))
	// fmt.Println(reconstructQueue1(data2))
}
