package pos

import "strconv"

func findMissingRanges(nums []int, lower int, upper int) []string {
	res := make([]string, 0)
	last := lower - 1
	nums = append(nums, upper+1)
	for i := 0; i < len(nums); i++ {
		if nums[i] == last+2 {
			res = append(res, strconv.Itoa(last+1))
		} else if nums[i] > last+2 {
			res = append(res, strconv.Itoa(last+1)+"->"+strconv.Itoa(nums[i]-1))
		}
		last = nums[i]
	}
	return res
}
