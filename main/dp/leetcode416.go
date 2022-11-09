package dp

func canPartition(nums []int) bool {
	// dp[i][j] using num[i] and before for sum j
	// dp[i][j] = dp[i-1][j] ｜｜ dp[i][j-nums[i]]
	sum := 0
	for _, value := range nums {
		sum = sum + value
	}
	if sum%2 == 1 || len(nums) == 1 {
		return false
	}
	sum = sum / 2
	dp := make([][]bool, len(nums))
	for index, _ := range dp {
		dp[index] = make([]bool, sum+1)
		dp[index][0] = true
	}
	for i := 1; i < len(nums); i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i] >= 0 {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)-1][sum]
}
