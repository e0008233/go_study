package stack_and_queue

type MyQueue struct {
	inStack  MyStack2
	outStack MyStack2
}

func Constructor1() MyQueue {
	return MyQueue{
		inStack:  ConstructorStack(),
		outStack: ConstructorStack(),
	}
}

func (this *MyQueue) Push(x int) {
	this.inStack.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.outStack.Empty() {
		for !this.inStack.Empty() {
			this.outStack.Push(this.inStack.Pop())
		}
	}
	return this.outStack.Pop()
}

func (this *MyQueue) Peek() int {
	if this.outStack.Empty() {
		for !this.inStack.Empty() {
			this.outStack.Push(this.inStack.Pop())
		}
	}
	return this.outStack.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.outStack.Empty() && this.inStack.Empty()
}

type MyStack2 struct {
	element []int
}

func ConstructorStack() MyStack2 {
	return MyStack2{
		element: make([]int, 0),
	}
}

func (this *MyStack2) Push(x int) {
	this.element = append(this.element, x)
}

func (this *MyStack2) Pop() int {
	value := this.Peek()
	this.element = this.element[0 : len(this.element)-1]
	return value
}

func (this *MyStack2) Peek() int {
	if this.Empty() {
		return 0
	}
	return this.element[len(this.element)-1]
}

func (this *MyStack2) Empty() bool {
	return len(this.element) == 0
}
