package pos

// 动态规划
func trap(height []int) int {
	res := 0
	maxRightHeight, maxLeftHeight := make([]int, len(height)), make([]int, len(height))
	for i := 1; i < len(height); i++ {
		maxLeftHeight[i] = max(height[i-1], maxLeftHeight[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		maxRightHeight[i] = max(height[i+1], maxRightHeight[i+1])
	}
	for i := range height {
		h := min(maxLeftHeight[i], maxRightHeight[i]) - height[i]
		if h > 0 {
			res += h
		}
	}
	return res
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 单调栈

// func trap(height []int) int{
// 	res := 0
// 	stack := make([]int, 0)
// 	for cur:= range height{
// 		for len(stack)>0 &&
// 			height[cur]>height[stack[len(stack)-1]]{
// 			popHeight := height[stack[len(stack)-1]]
// 			stack = stack[:len(stack)-1]
// 			if len(stack)==0{
// 				break
// 			}
// 			distance := cur-stack[len(stack)-1]-1
// 			boundedHeight := min(height[cur], height[stack[len(stack)-1]]) - popHeight
// 			res += distance*boundedHeight
//
// 		}
// 		stack = append(stack, cur)
// 	}
// 	return res
// }
