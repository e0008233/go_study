package greedy

import "math"

func MaxSubArray(nums []int) int {
	result := 0

	largest := math.MinInt
	for _, value := range nums {
		if result <= 0 && value <= 0 {
			if value >= largest {
				largest = value
			}
		} else if value >= 0 {
			result = result + value
			if result >= largest {
				largest = result
			}
		} else if value < 0 && value+result >= 0 {
			result = result + value
		} else {
			result = 0
		}
	}

	return largest
}
