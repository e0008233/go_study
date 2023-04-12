package arrays

func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	totalLength := len(matrix) * len(matrix[0])
	curr := totalLength / 2
	start := 0
	end := totalLength - 1
	i, j := getMatrixIndex(curr, len(matrix[0]))
	if matrix[i][j] == target {
		return true
	}

	for curr <= end && curr >= start {
		i, j = getMatrixIndex(curr, len(matrix[0]))

		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			start = curr + 1
			curr = (start + end) / 2
		} else {
			end = curr - 1
			curr = (start + end) / 2
		}
		i, j = getMatrixIndex(curr, len(matrix[0]))
		if matrix[i][j] == target {
			return true
		}

	}

	return false
}

// 0 1 2
// 3 4 5
func getMatrixIndex(index, column int) (int, int) {
	rowNum := index / column
	columNum := index - rowNum*column
	return rowNum, columNum
}
