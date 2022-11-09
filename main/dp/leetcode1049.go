package dp

func lastStoneWeightII(stones []int) int {

	// dp[i][j] := max(dp[i-1][j], dp[i-1][j-stones[i]]+stones[i])
	total := 0
	for _, value := range stones {
		total = total + value
	}
	sum := total / 2
	dp := make([][]int, len(stones))

	for index, _ := range dp {
		dp[index] = make([]int, sum+1)
		dp[index][0] = 0
	}
	for i := stones[0]; i <= sum; i++ {
		dp[0][i] = stones[0]
	}
	for i := 1; i < len(stones); i++ {
		for j := 1; j <= sum; j++ {
			if j-stones[i] >= 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-stones[i]]+stones[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return total - dp[len(stones)-1][sum] - dp[len(stones)-1][sum]
}
