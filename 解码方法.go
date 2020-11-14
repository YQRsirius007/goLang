package pos

func numDecodings(s string) int {
	// dp := make([]int, len(s)+1)
	if s[0] == '0' {
		return 0
	}
	prev, cur, tmp := 1, 1, 0 // dp[-1] dp[0]
	for i := 1; i < len(s); i++ {
		tmp = cur
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-2] == '2' {
				cur = prev
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			cur += prev
		}
		prev = tmp
	}
	return cur
}
