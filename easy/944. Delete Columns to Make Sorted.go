package easy

func minDeletionSize(strs []string) int {
	m, n := len(strs), len(strs[0])
	ans := 0
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			if strs[j][i] < strs[j-1][i] {
				ans += 1
				break
			}
		}
	}
	return ans
}
