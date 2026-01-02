package medium

func combinationSum260102(candidates []int, target int) [][]int {
	dp := make([][][]int, target+1)
	for _, candidate := range candidates {
		if candidate > target {
			continue
		}
		dp[candidate] = append(dp[candidate], []int{candidate})
		for i := candidate + 1; i <= target; i++ {
			if len(dp[i-candidate]) != 0 {
				for j := 0; j < len(dp[i-candidate]); j++ {
					tmp := make([]int, len(dp[i-candidate][j]))
					copy(tmp, dp[i-candidate][j])
					tmp = append(tmp, candidate)
					dp[i] = append(dp[i], tmp)
				}
			}
		}
	}
	return dp[target]
}
