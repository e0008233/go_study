package amazon

import "math"

func findMaxProducts(products []int32) int64 {
	// Write your code here
	sum := 0
	currentMax := 0
	for start := 0; start < len(products); start++ {
		for end := start; end < len(products); end++ {
			sum = findSum(start, end, products)
			if sum > currentMax {
				currentMax = sum
			}
		}
	}

	return int64(currentMax)
}

func findSum(start int, end int, products []int32) int {
	if start == end {
		return int(products[start])
	}
	min := math.MaxInt
	minIndex := 0
	for i := start; i <= end; i++ {
		if int(products[i]) < min {
			min = int(products[i])
			minIndex = i
		}
	}
	left := findMinMaxSum(start, minIndex-1, products, 0, min)
	right := findMinMaxSum(minIndex+1, end, products, min, math.MaxInt)
	if left == -1 || right == -1 {
		return -1
	} else {
		return left + right + min
	}
}

func findMinMaxSum(start int, end int, products []int32, min int, max int) int {
	if start == end {
		if int(products[start]) > min && int(products[start]) < max {
			return int(products[start])
		} else {
			return -1
		}
	}

	if start > end {
		return 0
	}

	min := math.MaxInt
	minIndex := 0
	for i := start; i <= end; i++ {
		if int(products[i]) < min {
			min = int(products[i])
			minIndex = i
		}
	}
	left := findMinMaxSum(start, minIndex-1, products, 0, min)
	right := findMinMaxSum(minIndex+1, end, products, min, math.MaxInt)
	if left == -1 || right == -1 {
		return -1
	} else {
		return left + right + min
	}
}
