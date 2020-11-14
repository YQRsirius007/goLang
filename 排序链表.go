package pos

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	midNode := middleNode(head)
	rightHead := midNode.Next
	midNode.Next = nil

	left := sortList(head)
	right := sortList(rightHead)
	return mergeTwoLists(left, right)
}

// 找中间节点，两个指针，快指针一次走两格，慢指针一格
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	sentry := &ListNode{Val: 0}
	cur := sentry

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return sentry.Next
}
