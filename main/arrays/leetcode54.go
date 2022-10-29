package arrays

type direction int

const (
	right direction = 0
	down  direction = 1
	left  direction = 2
	up    direction = 3
)

func spiralOrder(matrix [][]int) []int {
	rowLength := len(matrix)
	columnLength := len(matrix[0])
	endIndex := rowLength * columnLength
	result := make([]int, endIndex)
	isVisited := make([][]bool, rowLength)
	for index, _ := range isVisited {
		isVisited[index] = make([]bool, columnLength)
	}
	currDirection := right
	index := 0
	rowNum := 0
	columnNum := 0
	for {
		switch currDirection {
		case right:
			for {
				result[index] = matrix[rowNum][columnNum]
				isVisited[rowNum][columnNum] = true
				index++
				if columnNum+1 >= columnLength || isVisited[rowNum][columnNum+1] {
					rowNum++
					break
				}
				columnNum++
			}
		case down:
			for {
				result[index] = matrix[rowNum][columnNum]
				isVisited[rowNum][columnNum] = true
				index++
				if rowNum+1 >= rowLength || isVisited[rowNum+1][columnNum] {
					columnNum--
					break
				}
				rowNum++
			}
		case left:
			for {
				result[index] = matrix[rowNum][columnNum]
				isVisited[rowNum][columnNum] = true
				index++
				if columnNum-1 < 0 || isVisited[rowNum][columnNum-1] {
					rowNum--
					break
				}
				columnNum--
			}
		case up:
			for {
				result[index] = matrix[rowNum][columnNum]
				isVisited[rowNum][columnNum] = true
				index++
				if rowNum-1 < 0 || isVisited[rowNum-1][columnNum] {
					columnNum++
					break
				}
				rowNum--
			}

		}
		currDirection = (currDirection + 1) % 4
		if index >= endIndex {
			break
		}
	}
	return result
}
