package medium

func mostPoints(questions [][]int) int64 {
	dp := make([]int, len(questions))
	dp[len(questions)-1] = questions[len(questions)-1][0]
	for i := len(questions) - 2; i >= 0; i-- {
		selected := 0
		if i+questions[i][1]+1 >= len(questions) {
			selected = questions[i][0]
		} else {
			selected = questions[i][0] + dp[i+questions[i][1]+1]
		}
		dp[i] = max(dp[i+1], selected)
	}
	return int64(dp[0])
}
