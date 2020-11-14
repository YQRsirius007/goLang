package pos

// 解法1
func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	item := make([]int, 0)
	result = append(result, item)
	res_set := make(map[string]bool)
	generateWithDup(0, nums, &item, &result, &res_set)
	return result
}
func generateWithDup(i int, nums []int, item *[]int, result *[][]int, res_set *map[string]bool) {
	if i >= len(nums) {
		return
	}
	*item = append(*item, nums[i])
	// if !*res_set["123"]{
	// 	*result = append(*result, *item)
	// 	res_set[item] == true
	// }

	generate(i+1, nums, item, result)
	itemArr := *item
	*item = append([]int{}, itemArr[:len(*item)-1]...)
	generate(i+1, nums, item, result)
}

// 解法2
// func subsets2(nums []int) [][]int{
// 	n := len(nums)
// 	all_set := 1<<n
// 	result := make([][]int, 0)
// 	// item := make()
// 	for i:=0; i<all_set; i++{
// 		item := make([]int, 0)
// 		for j:=0; j<n; j++{
// 			if i&(1<<j)>0{
// 				item = append(item, nums[j])
// 				fmt.Println(j, item, nums, nums[j])
// 			}
// 		}
// 		result = append(result, item)
// 	}
// 	return result
// }
