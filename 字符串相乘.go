package pos

// 1、计算的方法
// 2、乘积的结果和对应结果的位置
// 3、特殊情况的处理 0*0
func multiply(num1 string, num2 string) string {
	// i, j := len(num1), len(num2)
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			mul := int((num1[i] - '0') * (num2[j] - '0'))
			p1, p2 := i+j, i+j+1
			sum := res[p2] + mul
			res[p1] += sum / 10
			res[p2] = sum % 10
		}
	}
	i := 0
	for i < len(res) && res[i] == 0 {
		i++
	}

	str := ""
	for ; i < len(res); i++ {
		str += string(res[i] + '0')
	}
	if str == "" {
		return "0"
	}
	return str
}

// 找到对应关系，无需反转
func reverse(s string) string {
	runes := []rune(s)
	for from, to := 0, len(s); from < to; from, to = from+1, to-1 {
		runes[to], runes[from] = runes[from], runes[to]
	}
	return string(runes)
}
