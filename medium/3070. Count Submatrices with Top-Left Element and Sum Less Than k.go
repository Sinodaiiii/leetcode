package medium

func countSubmatrices260318(grid [][]int, k int) int {
	n := len(grid[0])
	pre := make([]int, n)
	ans := 0
	for _, row := range grid {
		curr := make([]int, n)
		curr[0] = pre[0] + row[0]
		if curr[0] > k {
			break
		}
		ans += 1
		for i := 1; i < n; i++ {
			curr[i] = curr[i-1] + pre[i] - pre[i-1] + row[i]
			if curr[i] > k {
				break
			}
			ans += 1
		}
		pre = curr
	}
	return ans
}
