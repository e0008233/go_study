package greedy

func wiggleMaxLength(nums []int) int {
	// ignore the points on the trend
	if len(nums) <= 1 {
		return len(nums)
	}
	preDiff := 0

	result := 1

	for i := 1; i < len(nums); i++ {
		currDiff := nums[i] - nums[i-1]
		if currDiff == 0 {
			continue
		} else if currDiff > 0 && preDiff <= 0 {
			result++
		} else if currDiff < 0 && preDiff >= 0 {
			result++
		}

		preDiff = currDiff
	}

	return result
}
