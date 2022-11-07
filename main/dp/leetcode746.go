package dp

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	// 10,15,20
	// dp[n] = min (dp[n-1]+cost[n-1], dp[n-2]+cost[n-2])
	length := len(cost)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return 0
	}

	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= length; i++ {
		dp[i] = min(cost[i-2]+dp[i-2], cost[i-1]+dp[i-1])
	}
	return dp[length]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
