package dp

import "math"

func Jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		max := math.MaxInt
		for j := 0; j < i; j++ {

			if j+nums[j] >= i {
				if nums[j]+1 <= max {
					max = nums[j] + 1
				}
			}
		}
		dp[i] = max
	}
	return dp[len(nums)-1]
}
