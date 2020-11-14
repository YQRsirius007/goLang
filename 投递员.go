package pos

func postman(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	man := 0
	left, right := 0, 2
	for left < n && right < n {
		if arr[left] == 0 {
			left++
			right++
		} else {
			man++
			left = right + 1
			right = left + 2
		}
	}
	for left < n {
		if arr[left] == 1 {
			man++
			break
		}
		left++
	}
	return man
}
