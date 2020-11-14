package pos

// 动态规划 胜出!
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s))
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			res = max(res, dp[i])
		}
	}
	return res
}

// 栈
// func longestValidParentheses(s string)int{
// 	if len(s)==0{
// 		return 0
// 	}
// 	stack := make([]int, 0)
// 	stack = append(stack, -1)
// 	res := 0
// 	for i,v := range s{
// 		if v=='('{
// 			fmt.Println("(", i)
// 			stack = append(stack, i)
// 		}else{
// 			stack = stack[:len(stack)-1]
// 			if len(stack)==0{
// 				fmt.Println("分隔)", i)
// 				stack = append(stack, i)
// 			}else{
// 				fmt.Println(")", i, stack)
// 				res = max(res, i - stack[len(stack)-1])
// 			}
// 		}
// 	}
// 	return res
// }
