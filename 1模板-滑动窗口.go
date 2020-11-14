package pos

// 子串-》滑动窗口
// func minWindow(s string, t string)string{
// 	window := map[byte]int{}
// 	need := map[byte]int{}
// 	left, right := 0, 0
// 	windowStr := ""
// 	res := s
// 	for i := range t{
// 		need[i]++
// 	}
// 	for right<len(s){
// 		window = append(window, s[right])
// 		right++
// 		for ...{
// 			res = minLn(res, windowStr)
// 			windowStr = windowStr[:len(windowStr)-1]
// 			left++
// 		}
// 	}
// 	return res
// }
