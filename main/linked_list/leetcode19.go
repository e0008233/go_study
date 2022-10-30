package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fakeHead := &ListNode{
		Next: head,
	}
	fast := fakeHead
	slow := fakeHead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for {
		if fast == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return fakeHead.Next
}
