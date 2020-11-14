package pos

// partition切分操作总能排定一个元素 还能知道这个元素最终所在的位置
// 缩小范围
// 切分过程可以不借助额外的数组空间
// 用来查找-》类似二分查找
func findKthLargest(nums []int, k int) int {
	len := len(nums)
	left, right := 0, len-1
	target := len - k
	for left < right {
		index := partition1(nums, left, right)
		if index == target {
			return nums[index]
		}
		if index < target {
			left = index + 1
		} else {
			right = index - 1
		}
	}
	return 0
}
func partition1(arr []int, left int, right int) int {
	index := left
	key := arr[index]
	i, j := left, right
	for i < j {
		for i < j && arr[j] >= key {
			j--
		}
		if i < j {
			arr[i] = arr[j]
		}
		for i < j && arr[i] <= key {
			i++
		}
		if i < j {
			arr[j] = arr[i]
		}
	}
	arr[i] = key
	return i
}
func partition2(nums []int, left int, right int) int {
	key := nums[left]
	j := left
	for i := left + 1; i <= right; i++ {
		if nums[i] < key {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// [left+1, j)<key (j, i]>=key
	nums[j], nums[left] = nums[left], nums[j]
	return j
}
