package stack_and_queue

import "strconv"

func EvalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	stack := make([]int, 0)
	for _, v := range tokens {
		intVar, err := strconv.Atoi(v)
		if err == nil {
			stack = append(stack, intVar)
		} else {
			switch v {
			case "+":
				v1 := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				v2 := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]

				stack = append(stack, v1+v2)
			case "-":
				v1 := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				v2 := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]

				stack = append(stack, v2-v1)
			case "*":
				v1 := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				v2 := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]

				stack = append(stack, v1*v2)
			case "/":
				v1 := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				v2 := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]

				stack = append(stack, v2/v1)
			}
		}
	}

	return stack[0]
}
