package dp

func FindMaxForm(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs))
	for index, _ := range dp {
		dp[index] = make([][]int, m+1)
		for index2, _ := range dp[index] {
			dp[index][index2] = make([]int, n+1)
		}
	}

	for i := 0; i < len(dp); i++ {
		var zero, one int
		for _, v := range strs[i] {
			if v == '0' {
				zero++
			} else {
				one++
			}
		}
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				if i == 0 && j >= zero && k >= one {
					dp[i][j][k] = 1
				} else if i > 0 {
					if j-zero >= 0 && k-one >= 0 {
						dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-zero][k-one]+1)
					} else {
						dp[i][j][k] = dp[i-1][j][k]
					}

				}

			}
		}
	}
	return dp[len(strs)-1][m][n]
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
