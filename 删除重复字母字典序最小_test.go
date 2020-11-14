package pos

import (
	"testing"
)

func TestShouldReturnReceipt_remove(t *testing.T) {
	data := []string{
		"bcabc",
		"cbacdcbc",
		"baab",
		"bbcaac",
	}
	res := []string{
		"abc",
		"acdb",
		"ab",
		"bac",
	}
	for i := range data {
		c := removeDuplicateLetters(data[i])
		if c != res[i] {
			t.Errorf("want %s, but got %s", res[i], c)
		}
	}

}
