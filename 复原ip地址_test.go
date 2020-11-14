package pos

import (
	"testing"
)

func TestShouldReturnReceipt_ip(t *testing.T) {
	data := []string{
		"25525511135",
		"010010",
	}
	res := [][]string{
		{"255.255.11.135", "255.255.111.35"},
		{"0.10.0.10", "0.100.1.0"},
	}
	for i := range data {
		c := restoreIpAddresses(data[i])
		for j, v := range c {
			if v != res[i][j] {
				t.Errorf("want %v, but got %v", res[i], c)
			}
		}

	}
}
