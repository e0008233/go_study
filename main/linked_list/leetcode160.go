package linked_list

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	currA := headA
	currB := headB
	lengthA, lengthB := 0, 0
	for {
		currA = currA.Next
		lengthA++
		if currA == nil {
			break
		}
	}
	for {
		currB = currB.Next
		lengthB++
		if currB == nil {
			break
		}
	}
	var step int
	var fast, slow *ListNode
	if lengthB > lengthA {
		step = lengthB - lengthA
		fast = headB
		slow = headA
	} else {
		step = lengthA - lengthB
		fast = headA
		slow = headB
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}

	for {
		if fast == slow {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
