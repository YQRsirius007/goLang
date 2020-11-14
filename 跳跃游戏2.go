package pos

func jump(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	minStep := 0
	end, i := len(nums)-1, 0
	for end > 0 {
		i = 0
		for nums[i]+i < end {
			i++
		}
		end = i
		minStep++
	}
	return minStep
}
