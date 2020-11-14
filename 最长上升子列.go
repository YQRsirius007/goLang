package pos

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 1
	dp := make([]int, len(nums))
	for i := range nums {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
			if dp[i] > res {
				res = dp[i]
			}
		}
	}
	return res
}
