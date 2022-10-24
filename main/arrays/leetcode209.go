package arrays

import "math"

func MinSubArrayLen(target int, nums []int) int {
	leftPointer := 0
	result := math.MaxInt
	currentSum := 0
	for rightPointer, value := range nums {
		currentSum = currentSum + value
		if currentSum >= target {
			if rightPointer-leftPointer+1 < result {
				result = rightPointer - leftPointer + 1
			}
			for {
				currentSum = currentSum - nums[leftPointer]
				leftPointer++

				if currentSum >= target {
					if rightPointer-leftPointer+1 < result {
						result = rightPointer - leftPointer + 1
					}
				} else {
					break
				}
				if leftPointer > rightPointer {
					break
				}

			}

		}

	}

	if result != math.MaxInt {
		return result
	}
	return 0
}
