package stack_and_queue

func IsValid(s string) bool {
	stack := make([]int32, 0)

	for _, value := range s {
		switch value {
		case '(', '{', '[':
			stack = append(stack, value)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			val := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			switch value {
			case ')':
				if val != '(' {
					return false
				}
			case '}':
				if val != '{' {
					return false
				}
			case ']':
				if val != '[' {
					return false
				}
			}
		case ' ':

		default:
			return false
		}
	}
	return len(stack) == 0

}
