package dp

func numTrees(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum = sum + dp[j]*dp[i-j-1]
		}
		dp[i] = sum
	}
	return dp[n]
}
