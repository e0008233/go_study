package main

import (
	"fmt"
	"go_study/main/stack_and_queue"
)

func main() {

	fmt.Println(stack_and_queue.EvalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
