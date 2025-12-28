package easy

func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	i, j := m-1, 0
	ans := 0
	for i >= 0 && j < n {
		if grid[i][j] < 0 {
			ans += n - j
			i -= 1
		} else {
			j += 1
		}
	}
	return ans
}
