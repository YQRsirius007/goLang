package pos

func dandiaozhan(s string) {
	// 1 遍历
	ret := ""
	for i := 0; i < len(s); i++ {
		// 不满足单调性 回溯
		for len(ret) > 0 && s[i] > ret[len(ret)-1] {
			// 出栈
			ret = ret[:len(ret)-1]
		}
		// 满足递减性 入栈
		ret += string(s[i])
	}
	// 栈里的剩余元素 处理
}