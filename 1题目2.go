package pos

import (
	"fmt"
	"sort"
)

func numClosestPositions(positions [][]int, num int) [][]int {
	res := make([][]int, num)
	for i := range res {
		res[i] = make([]int, 2)
	}
	distance := make([][]int, len(positions))
	for i := range distance {
		distance[i] = make([]int, 2)
	}
	for i := range positions {
		area := positions[i][0]*positions[i][0] + positions[i][1]*positions[i][1]
		distance[i][0] = i
		distance[i][1] = area
	}
	sort.Slice(distance, func(i, j int) bool {
		return distance[i][1] < distance[j][1]
	})
	for i := 0; i < num; i++ {
		res[i] = positions[distance[i][0]]
	}
	return res
}

// [[9,0],[7,0],[1,9],[3,0],[2,7],[5,3],[6,0],[3,4],[6,2],[5,2]]
// [[3,0],[6,0],[7,0],[5,2],[5,3],[3,4],[9,0],[2,7],[6,2],[1,9]]
func reconstructQueue1(people [][]int) [][]int {
	if len(people) == 0 {
		return [][]int{}
	}
	// 体重排序
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] <= people[j][1]
		}
		return people[i][0] >= people[j][0]
	})
	res := make([][]int, len(people))
	for i := range res {
		res[i] = make([]int, 2)
	}
	for i := range people {
		idx := people[i][1]
		copy(res[idx+1:], res[idx:])
		res[idx] = people[i]
		fmt.Println(people[i], res)
	}
	return res
}
