package pos

import (
	"testing"
)

func TestShouldReturnReceipt_network(t *testing.T) {
	data := [][]int{
		{2, 1, 1},
		{2, 3, 1},
		{3, 4, 1},
	}
	c := networkDelayTime(data, 4, 2)
	if c != 2 {
		t.Errorf("want %d, but got %d", 2, c)
	}
}
