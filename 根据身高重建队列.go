package pos

import (
	"fmt"
	"sort"
)

func reconstrucQueue(people [][]int) [][]int {
	res := make([][]int, 0)
	if len(people) == 0 {
		return res
	}
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] <= people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	res = append(res, people[0])
	for i := 1; i < len(people); i++ {
		insert := people[i][1]
		fmt.Println(insert)
		copy(res[insert+1:], res[insert:])
		res[insert] = people[i]
		// res = append(res[:insert], append([][]int{people[i]}, res[insert:]...)...)
	}
	return res
}
func reconstructQueue3(people [][]int) [][]int {

	if len(people) == 0 {
		return people
	}
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] <= people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i+1])
		people[p[1]] = p

	}
	return people
}
