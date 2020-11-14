package pos

func findS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	r := make(map[int][]int)
	for k, v := range nums {
		if _, ok := r[v]; ok {
			r[v] = append(r[v], k)
		} else {
			a := make([]int, 1, 1)
			a[0] = k
			r[v] = a
		}
	}
	result := [2]int{0, 0}
	for _, v := range r {
		if result[0] < len(v) || result[0] == len(v) && result[1] > v[len(v)-1]-v[0] {
			result = [2]int{len(v), v[len(v)-1] - v[0]}
		}
	}
	return result[1] + 1
}
