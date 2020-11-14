package pos

import "strconv"

func addStrings(num1 string, num2 string) string {
	len1, len2, carry := len(num1), len(num2), 0
	res := ""
	i, j := len1-1, len2-1
	for i >= 0 || j >= 0 || carry != 0 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + carry
		res = strconv.Itoa(result%10) + res
		carry = result / 10
		i--
		j--
	}
	return res
}
