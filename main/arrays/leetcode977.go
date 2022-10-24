package arrays

func sortedSquares(nums []int) []int {
	leftPointer := 0
	rightPointer := len(nums) - 1
	result := make([]int, len(nums))
	index := len(nums) - 1
	for {
		left := nums[leftPointer] * nums[leftPointer]
		right := nums[rightPointer] * nums[rightPointer]
		if right > left {
			result[index] = right
			rightPointer--
			index--
		} else {
			result[index] = left
			leftPointer++
			index--
		}
		if leftPointer > rightPointer {
			break
		}
	}
	return result
}
