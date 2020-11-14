package pos

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	if nums == nil || len(nums) == 0 {
		return res
	}
	path := make([]int, 0)
	used := make([]bool, len(nums))
	dfs_permute(nums, len(nums), &path, &used, &res)
	return res
}

func dfs_permute(nums []int, lenRes int, path *[]int, used *[]bool, res *[][]int) {
	if len(*path) == lenRes {
		tmp := make([]int, lenRes)
		copy(tmp, *path)
		*res = append(*res, tmp)
		return
	}
	for i, v := range nums {
		if (*used)[i] {
			continue
		}
		*path = append(*path, v)
		(*used)[i] = true
		dfs_permute(nums, lenRes, path, used, res)
		*path = (*path)[:len(*path)-1]
		(*used)[i] = false
	}
}
