package pos

func dailyTemperatures(T []int) []int {
	length := len(T)
	ans := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temperature := T[i]
		for len(stack) > 0 && T[stack[len(stack)-1]] < temperature {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[preIndex] = i - preIndex
		}
		stack = append(stack, i)
	}
	return ans
}
