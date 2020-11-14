package pos

import "strconv"

func myAtoi(str string) int {
	INT_MAX := 1<<31 - 1
	INT_MIN := -1 << 31
	res := ""
	symbol := "+"
	for _, v := range str {
		if len(res) == 0 && len(symbol) == 1 && v == ' ' {
			continue
		}

		if len(res) == 0 && len(symbol) < 3 && v == '-' {
			symbol += "-"
		} else if len(res) == 0 && len(symbol) < 3 && v == '+' {
			symbol += "+"
		} else if v >= '0' && v <= '9' {
			res += string(v)
		} else {
			break
		}
	}
	if res == "" || len(symbol) > 2 {
		return 0
	}
	number, _ := strconv.Atoi(res)
	if symbol == "+" || symbol == "++" {
		if number > INT_MAX {
			return INT_MAX
		}
		return number
	} else {
		if -number < INT_MIN {
			return INT_MIN
		}
		return -number
	}
}
