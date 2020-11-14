package pos

import (
	"strconv"
)

func getPermutation(n int, k int) string {
	number := make([]string, 0)
	if n == 1 && k == 1 {
		return "1"
	}
	for i := 1; i <= n; i++ {
		numberStr := strconv.Itoa(i)
		number = append(number, numberStr)
	}
	res := ""
	factorial := make([]int, n+1)
	factorial[1] = 1
	for i := 2; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
	}
	for i := n; i > 1; i-- {
		base := factorial[i-1]
		curIndex := (k - 1) / base
		k = k % base
		if k == 0 {
			k = base
		}
		res += number[curIndex]
		number = append(number[:curIndex], number[curIndex+1:]...)
	}
	res += number[0]
	return res
}
