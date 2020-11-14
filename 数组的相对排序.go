package pos

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	res := []int{}
	sort.Ints(arr1)
	for p2 := 0; p2 < len(arr2); p2++ {
		c2 := arr2[p2]
		for p1 := 0; arr1[p1] <= c2; p1++ {
			if c2 == arr1[p1] {
				res = append(res, c2)
				arr1 = append(arr1[:p1], arr1[p1+1:]...)
			}
		}
	}
	for i := range arr1 {
		res = append(res, arr1[i])
	}
	return res
}
