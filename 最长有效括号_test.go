package pos

import "testing"

func TestShouldReturnReceipt_parentheses(t *testing.T) {
	data := []string{
		"()",
		"()())",
		"())()",
		")()(()))",
	}
	res := []int{
		2,
		4,
		2,
		6,
	}
	for i, v := range data {
		c := longestValidParentheses(v)
		if c != res[i] {
			t.Errorf("want %d, but got %d", res[i], c)
		}
	}
}
