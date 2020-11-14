package pos

// 输入：target = "10111"
// 输出：3
// 解释：初始配置 "00000".
// 从第 3 个灯泡（下标为 2）开始翻转 "00000" -> "00111"
// 从第 1 个灯泡（下标为 0）开始翻转 "00111" -> "11000"
// 从第 2 个灯泡（下标为 1）开始翻转 "11000" -> "10111"
// 至少需要翻转 3 次才能达成 target 描述的状态

func minFlips(target string) int {
	res := 0
	target = "0" + target
	for i := 1; i < len(target); i++ {
		if target[i] == 1 && target[i] != target[i-1] {
			res++
		}
	}
	return res
}
