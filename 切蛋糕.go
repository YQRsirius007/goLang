package pos

import (
	"sort"
)

func maxCake(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	maxH := max(horizontalCuts[0], h-horizontalCuts[len(horizontalCuts)-1])
	maxW := max(verticalCuts[0], w-verticalCuts[len(verticalCuts)-1])
	for i := 1; i < len(horizontalCuts); i++ {
		maxH = max(maxH, horizontalCuts[i]-horizontalCuts[i-1])
	}
	for i := 1; i < len(verticalCuts); i++ {
		maxW = max(maxW, verticalCuts[i]-verticalCuts[i-1])
	}
	return (maxH & 100000007) * (maxW & 100000007) % 1000000007
}
