package pos

import "strconv"

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	phone := []string{
		"",
		"!@#",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	path := ""
	dfs_letter(digits, 0, len(digits), phone, &path, &res)
	return res
}
func dfs_letter(digits string, index, digitsLen int, phone []string, path *string, res *[]string) {
	if len(*path) == digitsLen {
		*res = append(*res, *path)
	}
	if index >= digitsLen {
		return
	}
	number, _ := strconv.Atoi(digits[index : index+1])
	for i := 0; i < len(phone[number]); i++ {
		*path += phone[number][i : i+1]
		dfs_letter(digits, index+1, digitsLen, phone, path, res)
		*path = (*path)[:len(*path)-1]
	}
}

func calculate(s string) int {

	stack := make([]int, 0, len(s))
	op := make([]int, 0, len(s))
	num := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			num = 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			if len(op) > 0 && op[len(op)-1] > 2 {
				if op[len(op)-1] == 3 {
					stack[len(stack)-1] *= num
				} else {
					stack[len(stack)-1] /= num
				}
				op = op[:len(op)-1]
			} else {
				stack = append(stack, num)
			}
			// 退一格i 上面自动i++
			i--
		case '+':
			op = append(op, 1)
		case '-':
			op = append(op, -1)
		case '*':
			op = append(op, 3)
		case '/':
			op = append(op, 4)
		default:
			// fmt.Println(s[i])
		}
	}
	// fmt.Println(stack, op)
	for len(op) > 0 {
		stack[1] = stack[0] + op[0]*stack[1]
		op = op[1:]
		stack = stack[1:]
	}

	return stack[0]
}
