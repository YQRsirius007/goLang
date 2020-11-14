package pos

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	left, right, pos := head, head, head
	len := 1
	if head == nil {
		return nil
	}
	for pos.Next != nil {
		pos = pos.Next
		len++
	}
	k %= len
	if k == 0 {
		return head
	}
	for i := 0; i < k; i++ {
		right = right.Next
	}
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	res := left.Next
	right.Next = head
	left.Next = nil
	return res
}
