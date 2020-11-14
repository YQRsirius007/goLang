package pos

// 无限循环的弹出条件，识别到循环了且不为1-》快慢指针总会相等
func isHappy(n int) bool {
	slow, fast := n, n
	fast = squareSum(fast)
	fast = squareSum(fast)
	for slow != fast {
		slow = squareSum(slow)
		fast = squareSum(fast)
		fast = squareSum(fast)
	}
	if slow == 1 {
		return true
	}
	return false
}
func squareSum(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n = n / 10
	}
	return res
}
