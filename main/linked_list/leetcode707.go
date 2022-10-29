package linked_list

type MyLinkedList struct {
	head *Node
	size int
}

type Node struct {
	Val  int
	Next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
		size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index > this.size {
		return -1
	}
	currIndex := 0
	currNode := this.head
	for {
		if currIndex == index {
			if currNode != nil {
				return currNode.Val
			}
			return -1
		} else {
			currIndex++
			currNode = currNode.Next
		}
	}

}

func (this *MyLinkedList) AddAtHead(val int) {
	if this.head == nil {
		this.head = &Node{
			Val:  val,
			Next: nil,
		}
	} else {
		oldHead := this.head
		this.head = &Node{
			Val:  val,
			Next: oldHead,
		}
	}
	this.size = this.size + 1
}

func (this *MyLinkedList) AddAtTail(val int) {
	curr := this.head
	if curr == nil {
		this.head = &Node{
			Val:  val,
			Next: nil,
		}
	} else {
		for {
			if curr.Next == nil {
				break
			}
			curr = curr.Next
		}
		curr.Next = &Node{
			Val:  val,
			Next: nil,
		}
	}
	this.size = this.size + 1
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
	} else {
		currIndex := 0
		currNode := this.head
		for {
			if currIndex == index-1 {
				tempNode := currNode.Next
				currNode.Next = &Node{
					Val:  val,
					Next: tempNode,
				}
				break
			} else {
				currIndex++
				currNode = currNode.Next
			}
		}
	}
	this.size = this.size + 1
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}
	if index == 0 {
		this.head = this.head.Next
	} else {
		currIndex := 0
		currNode := this.head
		for {
			if currIndex == index-1 {
				currNode.Next = currNode.Next.Next
				break
			} else {
				currIndex++
				currNode = currNode.Next
			}
		}
	}
	this.size = this.size - 1
}
