package hot100

func combinationSum(candidates []int, target int) [][]int {
	// sort.Slice(candidates, func(i, j int) bool {return candidates[i] <= candidates[j]})
	n := len(candidates)
	dp := make([][][]int, target+1)
	for i := 0; i <= target; i++ {
		dp[i] = make([][]int, 0)
	}
	for i := 0; i < n; i++ {
		if candidates[i] <= target {
			dp[candidates[i]] = append(dp[candidates[i]], []int{candidates[i]})
		} else {
			continue
		}
		for j := candidates[i] + 1; j <= target; j++ {
			if len(dp[j-candidates[i]]) > 0 {
				for k := 0; k < len(dp[j-candidates[i]]); k++ {
					tmp := make([]int, len(dp[j-candidates[i]][k]))
					copy(tmp, dp[j-candidates[i]][k])
					tmp = append(tmp, candidates[i])
					dp[j] = append(dp[j], tmp)
				}
			}
		}
	}
	return dp[target]
}
