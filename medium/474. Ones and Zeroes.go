package medium

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	for _, str := range strs {
		currM, currN := 0, 0
		flag := true
		for i := 0; i < len(str); i++ {
			if str[i] == '0' {
				currM += 1
			} else if str[i] == '1' {
				currN += 1
			}
			if currM > m || currN > n {
				flag = false
				break
			}
		}
		if !flag {
			continue
		}
		for j := m; j >= currM; j-- {
			for k := n; k >= currN; k-- {
				if dp[j-currM][k-currN] >= 0 {
					dp[j][k] = max(dp[j][k], dp[j-currM][k-currN]+1)
				}
			}
		}
	}
	ans := 0
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			ans = max(ans, dp[i][j])
		}
	}
	return ans
}
