package bytedance

import "fmt"

func evaluateExp(s string) int32 {
	// Write your code here
	return evaluateExp2(s, 0, len(s)-1)

}

func evaluateExp2(s string, start, end int) int32 {
	fmt.Println("evaluateExp2", start, " ", end)
	// Write your code here
	if start == end {
		return int32(s[start] - '0')
	}
	var value int32 = -1
	expression := ""
	for i := start; i <= end; i++ {
		currentChar := string(s[i])

		switch currentChar {
		case "(":

			indexOfClosing := findClosingBracket(s, i, end)
			if indexOfClosing == -1 {
				fmt.Println("missing indexOfClosing")

				return -1
			}
			fmt.Println("( (", value, " ", expression)
			temp := evaluateExp2(s, i+1, indexOfClosing-1)
			if value != -1 && expression != "" {
				value = calculate(value, expression, temp)
				if value == -1 {
					fmt.Println("wrong calculate")

					return -1
				}
				expression = ""
			} else {
				value = temp
			}
			i = indexOfClosing
			continue
		case ")":
			fmt.Println("wrong )")
			return -1
		case "#", "$":
			if expression != "" {
				fmt.Println("wrong expression")

				return -1
			}
			expression = currentChar

		default:
			temp := int32(s[i] - '0')
			if temp >= 10 || temp < 0 {
				fmt.Println("wrong temp")
				return -1
			}
			if value != -1 && expression != "" {
				value = calculate(value, expression, temp)
				if value == -1 {
					fmt.Println("wrong calculate")
					return -1
				}
				expression = ""
			} else {
				value = temp
			}

		}
	}
	fmt.Println("final", value)
	return value
}

func calculate(value int32, expression string, temp int32) int32 {
	fmt.Println("calculate start %v", value, expression, temp)
	switch expression {
	case "#":
		return min(value, temp)
	case "$":
		return max(value, temp)
	default:
		return -1
	}
}

func max(value int32, temp int32) int32 {
	if value > temp {
		return value
	}
	return temp
}

func min(value int32, temp int32) int32 {
	if value < temp {
		return value
	}
	return temp
}

func findClosingBracket(s string, i, end int) int {
	bracketCount := 1
	for j := i + 1; j <= end; j++ {
		currentChar := string(s[j])
		if currentChar == ")" {
			bracketCount--
			if bracketCount == 0 {
				return j
			}
		} else if currentChar == "(" {
			bracketCount++
		}
	}
	return -1
}
