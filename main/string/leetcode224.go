package string

import (
	"go_study/main/util"
	"strconv"
)

func Calculate(s string) int {
	// https://www.youtube.com/watch?v=081AqOuasw0
	// Write your code here
	ss := []byte(s)
	result := 0
	var operator byte

	stack := util.NewStack()

	for i := 0; i < len(ss); i++ {
		char := ss[i]
		if char == ' ' {
			continue
		} else if char == '+' {
			operator = '+'
		} else if char == '-' {
			operator = '-'
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
				if oldOperator.(byte) == '+' {
					result = result + oldResult.(int)
				} else {
					result = oldResult.(int) - result
				}
				operator = 0
			}
		} else if char >= '0' && char <= '9' {
			num := 0
			for i < len(ss) {
				char = ss[i]
				if char >= '0' && char <= '9' {
					temp, _ := strconv.Atoi(string(char))
					num = num*10 + temp
					i++
				} else {
					break
				}
			}
			i--
			if operator != 0 {
				if operator == '+' {
					result = result + num
				} else {
					result = result - num
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
	return result
}
