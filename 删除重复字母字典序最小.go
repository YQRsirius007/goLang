package pos

import (
	"strings"
)

// 用栈处理，符合条件的加进去，遇到什么情况出栈
// 不该顺序的去除后面的重复字符
// 直接按顺序添加到结果中，后面有重复的不加就可以
func removeDuplicateLetters(s string) string {
	times := make([]int, 26)
	for _, c := range s {
		times[c-'a']++
	}
	res := ""
	for i := range s {
		if !strings.Contains(res, string(s[i])) {
			for len(res) > 0 && res[len(res)-1] > s[i] && times[res[len(res)-1]-'a'] > 0 {
				res = res[:len(res)-1]
			}
			res += string(s[i])
		}
		times[s[i]-'a']--
	}
	return res
}
func xx(s string) string {
	ret := ""
	return ret
}
