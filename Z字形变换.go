package pos

func convert(s string, numRows int) string {
	// z字型变换 在边界时刻， 箭头的方向发生变化
	// 最后返回一个string 中间过程也无需存0 直接存结果 好返回
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	curRow := 0
	flag := false
	for _, c := range s {
		rows[curRow] += string(c)
		if curRow == 0 || curRow == numRows-1 {
			flag = !flag
		}
		if flag {
			curRow++
		} else {
			curRow--
		}
	}
	res := ""
	for _, row := range rows {
		res += row
	}
	return res
}
