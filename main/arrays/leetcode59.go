package arrays

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for index, _ := range result {
		result[index] = make([]int, n)
	}
	count := 1
	currentDirection := right
	rowNum := 0
	columnNum := 0
	for {
		switch currentDirection {
		case right:
			for {
				result[rowNum][columnNum] = count
				count++
				if columnNum+1 >= n || result[rowNum][columnNum+1] != 0 {
					rowNum++
					break
				}
				columnNum++
			}
		case down:
			for {
				result[rowNum][columnNum] = count
				count++
				if rowNum+1 >= n || result[rowNum+1][columnNum] != 0 {
					columnNum--
					break
				}
				rowNum++
			}
		case left:
			for {
				result[rowNum][columnNum] = count
				count++
				if columnNum-1 < 0 || result[rowNum][columnNum-1] != 0 {
					rowNum--
					break
				}
				columnNum--
			}
		case up:
			for {
				result[rowNum][columnNum] = count
				count++
				if rowNum-1 < 0 || result[rowNum-1][columnNum] != 0 {
					columnNum++
					break
				}
				rowNum--
			}
		}
		currentDirection = (currentDirection + 1) % 4
		if count > n*n {
			break
		}
	}

	return result
}
