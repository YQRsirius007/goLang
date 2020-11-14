package pos

import (
	"fmt"
	"strconv"
	"strings"
)

// 入栈和出栈的操作
func decodeString(s string) string {
	stk := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else {
			// 出栈
			ptr++
			sub := []string{}
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			// 反转
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			// 把[取出来
			stk = stk[:len(stk)-1]
			// 取数字
			repTime, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1]
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t)
		}
	}
	return getString(stk)
}
func getDigits(s string, ptr *int) string {
	ret := ""
	for s[*ptr] >= '0' && s[*ptr] <= '9' {
		ret += string(s[*ptr])
		*ptr++
	}
	return ret
}
func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}

// a(b(c)<3>)<2>d

func decodeString2(s string) string {
	stackStr := []byte{}
	stackIdx := []int{}
	// ret := ""
	var curNum int
	for i := 0; i < len(s); i++ {
		// fmt.Println(string(s[i]), stackStr, stackIdx)
		if s[i] == '(' {
			stackIdx = append(stackIdx, len(stackStr))
		} else if s[i] == '<' {
			curNum = 0
		} else if s[i] == '>' {
			str := string(stackStr[stackIdx[len(stackIdx)-1]:])
			stackStr = stackStr[:stackIdx[len(stackIdx)-1]]
			fmt.Println(str, curNum)
			stackIdx = stackIdx[:len(stackIdx)-1]

			stackStr = append(stackStr, []byte(strings.Repeat(str, curNum))...)
		} else if s[i] >= '0' && s[i] <= '9' {
			curNum = curNum*10 + int(s[i]-'0')
		} else if s[i] == ')' {
			// stackStr = append(stackStr, s[i])
		} else {
			stackStr = append(stackStr, s[i])
		}
	}
	return string(stackStr)
}
func getNum(s string, ptr *int) int {
	num := 0
	fmt.Println("getNum", s[*ptr]-'0')
	for s[*ptr] >= '0' && s[*ptr] <= '9' {
		temp, _ := strconv.Atoi(string(s[*ptr]))
		num = num*10 + temp
		*ptr++
	}
	fmt.Println("getNum", num)
	return num
}
