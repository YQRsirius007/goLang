package pos

// 存对应关系的时候，可以用两个数组，不一定是map
func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""
	for i := 0; i < len(values) && num > 0; i++ {
		for values[i] <= num {
			res += symbol[i]
			num -= values[i]
		}
	}
	return res
}
