package pos

import "fmt"

func findShortestSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	left := map[int]int{}
	right := map[int]int{}
	count := map[int]int{}
	for i, v := range nums {
		if _, ok := left[v]; !ok {
			left[v] = i
		}
		right[v] = i
		count[v]++
	}
	degree := 1
	for _, v := range count {
		if v > degree {
			degree = v
		}
	}
	fmt.Println(degree, count)
	ans := len(nums)
	for _, v := range count {
		if v == degree && ans > right[v]-left[v]+1 {
			ans = right[v] - left[v] + 1
		}
	}
	return ans
}
