package pos

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ret := [][]int{}
	if nums == nil || len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	n := len(nums)
	// fmt.Println(nums, n)
	index, left, right := 0, 1, n-1
	for index < n-2 {

		if nums[index] > 0 {
			break
		}
		if index > 0 && nums[index] == nums[index-1] {
			index++
			continue
		}
		left = index + 1
		right = n - 1
		// fmt.Println(index, left, right)
		for left < right {
			fmt.Println(nums[index], nums[left], nums[right], nums[left]+nums[right]+nums[index])
			if nums[left]+nums[right]+nums[index] > 0 {
				right--
				continue
			}
			if nums[left]+nums[right]+nums[index] < 0 {
				left++
				continue
			}
			if left == index+1 || nums[left] != nums[left-1] {
				ret = append(ret, []int{nums[index], nums[left], nums[right]})
			}
			left++
			right--
		}
		index++
	}
	return ret
}
