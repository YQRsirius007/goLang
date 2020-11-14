package pos

import "fmt"

func isLongPressedName(name string, typed string) bool {
	if name == typed {
		return true
	}
	left, right := 0, 0
	for ; left < len(name) && right < len(typed); left++ {
		if typed[right] != name[left] {
			return false
		}
		right++
		if left+1 < len(name) && name[left] == name[left+1] {
			continue
		}
		for right < len(typed) && typed[right] == name[left] {
			right++
		}
	}
	fmt.Println(left, right)
	if left == len(name) && right == len(typed) {
		return true
	}
	return false
}
