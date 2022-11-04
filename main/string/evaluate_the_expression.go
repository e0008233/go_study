package string

import (
	"go_study/main/util"
	"strconv"
)

//Question Description
//You are given an expression s, it consists of numbers(0-9), parentheses, '#' and '$"
// $ means max operation: 3$4 = 4
// # means min operation: 3#4 = 3
// $ and # share the same precedence: 3$1#2 = 2
// Parentheses means higher precedence, it can be embedded: 2#(3$(1#2))
//
// Make it simple, the numbers in the expression are in [0,9]
// Your goal is to evaluate the expression, and return the final result. return -1 if the expression is invalid.
// Constraints
// 1 <= len(s)<= 1000

func EvaluateExp(s string) int32 {
	// Write your code here
	ss := []byte(s)
	result := 0
	var operator byte

	stack := util.NewStack()

	for _, char := range ss {
		if char == '#' {
			operator = '#'
		} else if char == '$' {
			operator = '$'
		} else if char == '(' {
			stack.Push(operator)
			stack.Push(result)
			result = 0
			operator = 0
		} else if char == ')' {
			oldResult := stack.Pop()
			oldOperator := stack.Pop()
			if oldResult == nil || oldOperator == nil {
				return -1
			}
			if oldOperator.(byte) != 0 {
				if oldOperator.(byte) == '#' {
					result = min(result, oldResult.(int))
				} else {
					result = max(result, oldResult.(int))
				}
				operator = 0
			}
		} else if char >= '0' && char <= '9' {
			num, _ := strconv.Atoi(string(char))
			if operator != 0 {
				if operator == '#' {
					result = min(result, num)
				} else {
					result = max(result, num)
				}
				operator = 0
			} else {
				result = num
			}
		}
	}
	if stack.Pop() != nil {
		return -1
	}
	return int32(result)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
