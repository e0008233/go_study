package stack_and_queue

type MyStack struct {
	element MyQueue2
}

func Constructor() MyStack {
	return MyStack{
		element: ConstructorMyQueue2(),
	}
}

func (this *MyStack) Push(x int) {
	this.element.Push(x)
}

func (this *MyStack) Pop() int {
	count := this.element.size()
	for count > 1 {
		count--
		this.element.Push(this.element.Dequeue())
	}
	return this.element.Dequeue()
}

func (this *MyStack) Top() int {
	val := this.Pop()
	this.element.Push(val)
	return val
}

func (this *MyStack) Empty() bool {
	return this.element.Empty()
}

type MyQueue2 struct {
	element []int
}

func ConstructorMyQueue2() MyQueue2 {
	return MyQueue2{
		element: make([]int, 0),
	}
}

func (this *MyQueue2) Push(x int) {
	this.element = append(this.element, x)
}

func (this *MyQueue2) Dequeue() int {
	if this.Empty() {
		return 0
	}
	value := this.Peek()
	this.element = this.element[1:len(this.element)]
	return value
}

func (this *MyQueue2) Peek() int {
	if this.Empty() {
		return 0
	}
	return this.element[0]
}

func (this *MyQueue2) Empty() bool {
	return len(this.element) == 0
}

func (this *MyQueue2) size() int {
	return len(this.element)
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
