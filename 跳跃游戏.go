package pos

func jump1(nums []int, n int) bool {
	// if nums==nil ||len(nums)==0 || n==0{
	// 	return false
	// }
	maxDistance := 0
	for i := 0; i < n; i++ {
		if i <= maxDistance {
			maxDistance = max(maxDistance, i+nums[i])
		}
		if maxDistance >= n-1 {
			return true
		}
	}
	return false
}
