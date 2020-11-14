package pos

import (
	"fmt"
	"testing"
)

func TestShouldReturnReceipt_removeKdigits(t *testing.T) {
	s := "abc"
	fmt.Println(s[:2])
	data := []string{
		// "1432219",
		"10200",
		"1234567890",
	}
	k := []int{
		// 3,
		1,
		9,
	}
	res := []string{
		// "1219",
		"200",
		"0",
	}
	for i, _ := range data {
		c := removeKdigits(data[i], k[i])
		if c != res[i] {
			t.Errorf("want %s, but got %s", res[i], c)
		}
	}
}
