package amazon

import "sort"

func minimalHeaviestSetA(arr []int32) []int32 {
	// Write your code here
	if len(arr) == 1 {
		return arr
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	var sum int32 = 0
	for _, value := range arr {
		sum = sum + value
	}
	half := sum / 2

	var result []int32

	var currentSum int32 = 0
	for _, value := range arr {
		currentSum = currentSum + value
		result = append(result, value)
		if currentSum > half {
			sort.Slice(result, func(i, j int) bool {
				return result[i] < result[j]
			})
			return result
		}
	}
	return result
}
