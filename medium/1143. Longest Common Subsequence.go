package medium

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m)
	find := false
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if text1[i] == text2[0] {
			find = true
		}
		if find {
			dp[i][0] = 1
		}
	}
	find = false
	for i := 0; i < n; i++ {
		if text2[i] == text1[0] {
			find = true
		}
		if find {
			dp[0][i] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = max(dp[i-1][j-1]+1, dp[i-1][j], dp[i][j-1])
			} else {
				dp[i][j] = max(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m-1][n-1]
}
