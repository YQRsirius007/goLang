package pos

import (
	"fmt"
	"testing"
)

func TestShouldReturnReceipt_timu1(t *testing.T) {
	s := "abc"
	fmt.Println(s[:1])
	data := []string{
		"23rd Oct 2020",
		"",
		"",
	}
	res := []string{
		"2020-10-23",
		"",
		"",
	}
	for i, v := range data {
		c := changeDateFormat(v)
		if c != res[i] {
			t.Errorf("want %s, but got %s", res[i], c)
		}
	}
}
