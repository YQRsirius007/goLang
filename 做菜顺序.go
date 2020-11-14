package pos

import (
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] > satisfaction[j]
	})

	if len(satisfaction) == 0 || satisfaction[0] <= 0 {
		return 0
	}
	res, sum := 0, 0
	for _, item := range satisfaction {
		sum += item
		if sum < 0 {
			break
		}
		res += sum
	}
	return res
}
