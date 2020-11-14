package pos

import "fmt"

func drink(n int) int {
	res, k, m := 0, 0, 0
	fmt.Println(n, res)
	if n < 3 {
		return 0
	}
	for n > 2 {
		k = n / 3
		m = n % 3
		res += k
		n = m + k
	}
	if n == 2 {
		res += 1
	}
	fmt.Println(n, res)
	return res
}
