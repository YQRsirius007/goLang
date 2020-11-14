package pos

import (
	"fmt"
	"runtime"
	"testing"
)

func TestShouldReturnReceipt_qqq(t *testing.T) {
	data := []int{
		3,
		10,
		81,
		59,
		0,
		94,
	}
	res := []int{
		1,
		5,
		40,
		29,
		0,
		47,
	}
	fmt.Println("cpu", runtime.NumCPU())
	for i, v := range data {
		c := drink(v)
		if c != res[i] {
			t.Errorf("want %d, but got %d", res[i], c)

		}
	}
}
