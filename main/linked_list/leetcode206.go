package linked_list

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	curr := head.Next
	prev.Next = nil

	for {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
		if curr == nil {
			break
		}
	}
	return prev
}
