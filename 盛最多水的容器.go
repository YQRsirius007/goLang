package pos

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	maxArea := 0
	left, right := 0, len(height)-1
	for left < right {
		heightLeft, heightRight := height[left], height[right]
		area := (right - left) * min(heightLeft, heightRight)
		maxArea = max(maxArea, area)
		if heightLeft < heightRight {
			for height[left] < heightLeft {
				left++
			}
		} else {
			for height[right] < heightRight {
				right++
			}
		}
	}
	return maxArea
}
