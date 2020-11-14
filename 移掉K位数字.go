package pos

// 保持原来的位置不变
// 相同位数 取决第一个不同数字的大小
// 遍历 保留遍历到的元素 选择性丢弃之前的元素
func removeKdigits(num string, k int) string {
	if len(num) == 0 {
		return ""
	}
	if k == 0 {
		return num
	}
	if k == len(num) {
		return "0"
	}
	remain := len(num) - k
	res := ""
	for i := range num {
		// 判断是否丢弃 结束条件
		for len(res) > 0 && k > 0 && num[i] < res[len(res)-1] {
			res = res[:len(res)-1]
			k--
		}
		res += string(num[i])
	}
	// fmt.Println(res)
	// 结果去掉前导零
	for len(res) > 1 && res[0] == '0' {
		res = res[1:]
	}
	if len(res) > remain {
		return res[:remain]
	}
	return res
}
