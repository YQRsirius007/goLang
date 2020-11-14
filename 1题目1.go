package pos

import (
	"fmt"
	"strings"
)

func changeDateFormat(date string) string {
	ret := ""
	arr := strings.Fields(date)
	if len(arr) < 3 {
		return ret
	}
	fmt.Println(date, arr[2])
	ret = year(arr[2])
	ret += "-"
	ret += month[arr[1]]
	ret += "-"
	ret += day(arr[0])
	return ret
}
func day(s string) string {
	if len(s) == 3 {
		return "0" + s[:1]
	}
	return s[:2]
}
// switch case圈复杂度过高 换成map
var month = map[string]string{
	"Jan":"01",
}
// func month(s string) string {
// 	switch s {
// 	case "Jan":
// 		return "01"
// 	case "Feb":
// 		return "02"
// 	case "Mar":
// 		return "03"
// 	case "Apr":
// 		return "04"
// 	case "May":
// 		return "05"
// 	case "Jun":
// 		return "06"
// 	case "Jul":
// 		return "07"
// 	case "Aug":
// 		return "08"
// 	case "Sep":
// 		return "09"
// 	case "Oct":
// 		return "10"
// 	case "Nov":
// 		return "11"
// 	case "Dec":
// 		return "12"
// 	default:
// 		return ""
// 	}
// }
func year(s string) string {
	return s
}

func removeDuplicateLetters1(s string) string {
	idx := make([]int, 26)
	for i := range idx {
		idx[i] = -1
	}
	ret := ""
	// 保证不重复
	isStack := [256]bool{}
	// 出栈条件判断：之后是否还有
	count := [256]int{}
	for i := range s {
		count[s[i]]++
	}
	for i := range s {
		count[s[i]]--
		if !isStack[s[i]] {
			fmt.Println(ret, string(s[i]), count[s[i]])
			for j := len(ret) - 1; j >= 0 && ret[j] > s[i]; j-- {
				fmt.Println("ready", string(ret[j]))
				if count[ret[j]] == 0 {
					break
				}
				// 出栈
				isStack[ret[len(ret)-1]] = false
				ret = ret[:len(ret)-1]
			}
			ret = ret + string(s[i])
			isStack[s[i]] = true
		}
	}
	return ret
}
