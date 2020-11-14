package pos

func canCross(stones []int) bool {
	if len(stones) == 1 {
		return true
	}
	if stones[1] != 1 {
		return false
	}
	l := len(stones)
	dp := make([][]bool, l)
	for i := range dp {
		dp[i] = make([]bool, l+1)
	}
	dp[0][0] = true
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			k := stones[i] - stones[j]
			if k <= i {
				dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
				if i == l-1 && dp[i][k] {
					return true
				}
			}
		}
	}
	return false
}
