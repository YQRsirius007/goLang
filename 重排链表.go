package pos

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	right := head
	var list []*ListNode
	for right != nil {
		list = append(list, right)
		right = right.Next
	}
	if len(list) == 2 {
		return
	}
	// p := left
	// fmt.Println(len, list)
	slow := 0
	fast := len(list)
	for slow < fast {
		list[slow].Next = list[fast]
		slow++
		list[fast].Next = list[slow]
		fast--
	}
	list[slow].Next = nil
	return
}
