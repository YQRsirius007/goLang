package pos

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	n := len(stations)
	dp := make([]int, n+1)
	dp[0] = startFuel
	for i := 0; i < n; i++ {
		for t := i; t > 0; t++ {
			if stations[i][0] <= dp[t-1] {
				dp[t] = max(dp[t], dp[t-1]+stations[i][1])
			}
		}
	}
	for i := range dp {
		if dp[i] >= target {
			return i
		}
	}
	return -1
}
