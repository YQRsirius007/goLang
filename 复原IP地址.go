package pos

import (
	"fmt"
	"strconv"
	"strings"
)

// 剪枝条件
// 递归中止条件
// 每个节点代表求解的不同的阶段，需要的状态变量
// splitTimes begin path res
func restoreIpAddresses(s string) []string {
	len := len(s)
	res := make([]string, 0)
	if len < 4 || len > 12 {
		return res
	}
	path := make([]string, 0)
	splitTimes := 0
	dfsIp(s, len, splitTimes, 0, &path, &res)
	return res
}

// 判断s的子区间[left, right]是否能够成为一个ip段
// 判断的同时顺便把类型转了
func isIpSegement(s string, left, right int) bool {
	len := right - left + 1
	if len > 1 && s[left:left+1] == "0" {
		return false
	}
	num, _ := strconv.Atoi(s[left : right+1])
	fmt.Println(num)
	if num > 255 {
		return false
	}
	return true
}
func dfsIp(s string, sLen, split, begin int, path, res *[]string) {
	// 递归中止
	if begin == sLen {
		if split == 4 {
			*res = append(*res, strings.Join(*path, "."))
		}
		return
	}
	// 剪枝
	if sLen-begin < (4-split) || sLen-begin > 3*(4-split) {
		return
	}

	for i := 0; i < 3; i++ {
		if begin+i >= sLen {
			break
		}
		ipSegment := isIpSegement(s, begin, begin+i)
		if ipSegment {
			*path = append(*path, s[begin:begin+i+1])
			dfsIp(s, sLen, split+1, begin+i+1, path, res)
			*path = (*path)[:len(*path)-1]
		}
	}
}
