package pos

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	if len(candidates) == 0 {
		return res
	}
	path := make([]int, 0)
	sum := 0
	dfs_combinationSum(candidates, target, &res, &path, &sum, 0)
	return res
}
func dfs_combinationSum(candidates []int, target int, res *[][]int, path *[]int, sum *int, index int) {
	if *sum == target {
		temp := make([]int, len(*path))
		copy(temp, *path)
		*res = append(*res, temp)
		return
	}
	if *sum > target {
		return
	}
	for i := index; i < len(candidates); i++ {
		*path = append(*path, candidates[i])
		*sum += candidates[i]
		dfs_combinationSum(candidates, target, res, path, sum, i)
		*path = (*path)[:len(*path)-1]
		*sum -= candidates[i]
	}
}
