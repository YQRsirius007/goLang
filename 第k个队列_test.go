package pos

import (
	"testing"
)

func TestShouldReturnReceipt_k(t *testing.T) {
	dataN := []int{
		3,
		3,
		3,
		4,
		4,
	}
	dataK := []int{
		3,
		2,
		6,
		9,
		8,
	}
	res := []string{
		"213",
		"132",
		"321",
		"2314",
		"23",
	}
	for i, v := range dataN {
		c := getPermutation(v, dataK[i])
		if c != res[i] {
			t.Errorf("want %s, but got %s", res[i], c)
		}
	}
}
