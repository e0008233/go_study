package linked_list

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	newHead := head
	prev := head
	curr := head

	for {
		if curr.Val == val {
			if curr == newHead {
				curr = curr.Next
				newHead = newHead.Next
				prev = prev.Next
			} else {
				prev.Next = curr.Next
				curr = curr.Next
			}
		} else {
			prev = curr
			curr = curr.Next
		}

		if curr == nil {
			break
		}
	}

	return newHead

}
