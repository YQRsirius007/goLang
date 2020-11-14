package pos

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	dfs_thesis(&res, n, n, "")
	return res
}
func dfs_thesis(res *[]string, left, right int, path string) {
	// 结束 return
	if left == 0 && right == 0 {
		*res = append(*res, path)
		return
	}
	// 剪枝,一个一个往后加的有加了一个没有左括号配对的右括号
	if right < left {
		return
	}
	if left > 0 {
		dfs_thesis(res, left-1, right, path+"(")
	}
	if right > 0 {
		dfs_thesis(res, left, right, path+")")
	}
}
