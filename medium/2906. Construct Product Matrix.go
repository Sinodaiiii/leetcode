package medium

func constructProductMatrix260522(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}
	pre := 1
	ans[0][0] = 1
	for x := 1; x < m*n; x++ {
		pre = (pre * grid[(x-1)/n][(x-1)%n]) % 12345
		ans[x/n][x%n] = pre
	}
	pre = 1
	for x := m*n - 2; x >= 0; x-- {
		pre = (pre * grid[(x+1)/n][(x+1)%n]) % 12345
		ans[x/n][x%n] = (pre * ans[x/n][x%n]) % 12345
	}
	// ans[m-1][n-1] = ans[m-1][n-1] % 12345
	return ans
}
