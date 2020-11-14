package pos

import (
	"fmt"
	"sort"
)

func bestTeamScore(scores []int, ages []int) int {
	if len(scores) == 0 || len(ages) == 0 {
		return 0
	}
	// 不能直接根据别人的数组去搞  二维数组
	order := make([][2]int, len(scores))
	for i := range order {
		order[i] = [2]int{ages[i], scores[i]}
	}
	sort.Slice(order, func(i, j int) bool {
		return order[i][0] < order[j][0] || (order[i][0] == order[j][0] && order[i][1] < order[j][1])
	})
	fmt.Println(order)

	ret := 0
	dp := make([]int, len(scores))
	for i := range order {
		// idx := order[i]
		dp[i] = order[i][1]
		cur := 0
		for j := 0; j < i; j++ {
			if dp[i] >= order[j][1] {
				cur = max(cur, dp[j])
			}
		}
		dp[i] += cur
		ret = max(ret, dp[i])
	}
	return ret
}
func dfs_score(idx int, scores []int, ages []int, choosed *[]bool, sum int, res *[]int) {
	if idx == len(scores) {
		*res = append(*res, sum)
		return
	}
	for i := 0; i < idx; i++ {

	}
}
