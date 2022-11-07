package dp

func IntegerBreak(n int) int {
	dp := make([]int, n+1)

	dp[1] = 1
	for i := 2; i <= n; i++ {
		currentMax := 0
		for j := 1; j < i; j++ {
			currentMax = max(max(j*(i-j), j*dp[i-j]), currentMax)
			product := j * dp[i-j]
			if product > currentMax {
				currentMax = product
			}
		}
		dp[i] = currentMax
	}

	return dp[n]
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
