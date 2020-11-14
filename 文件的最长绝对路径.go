package pos

import (
	"fmt"
	"strings"
)

func lengthLongestPath(input string) int {
	if len(input) == 0 {
		return 0
	}
	stack := make([]int, len(input)+1)
	res := 0
	arr := strings.Split(input, "\n")
	for _, s := range arr {
		level := strings.Count(s, "\t") + 1
		len := len(s) - (level - 1)
		fmt.Println(level, len, 2*strings.Count(s, "\t"))
		if strings.Contains(s, ".") {
			res = max(res, stack[level-1]+len)
		} else {
			stack[level] = stack[level-1] + len + 1
		}
	}
	return res
}
