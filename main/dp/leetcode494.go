package dp

func findTargetSumWays(nums []int, target int) int {
	return helper(nums, 0, target, 0, 0)
}

func helper(nums []int, index, target, sum, result int) int {
	if index == len(nums)-1 {
		if sum+nums[index] == target {
			result = result + 1
		}
		if sum-nums[index] == target {
			result = result + 1
		}
		return result
	}

	result1 := helper(nums, index+1, target, sum+nums[index], result)
	result2 := helper(nums, index+1, target, sum-nums[index], result)
	return result1 + result2
}
