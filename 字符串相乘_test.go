package pos

import "testing"

func TestShouldReturnReceipt_multiply(t *testing.T) {
	data1 := []string{
		"2",
		"123",
		"0",
	}
	data2 := []string{
		"3",
		"456",
		"0",
	}
	res := []string{
		"6",
		"56088",
		"0",
	}
	for i, v := range data1 {
		c := multiply(v, data2[i])
		if c != res[i] {
			t.Errorf("want %s, but got %s", res[i], c)
		}
	}
}
