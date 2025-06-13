package hot100

func numIslands(grid [][]byte) int {
	ans := 0
	n, m := len(grid), len(grid[0])
	direct := [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	var detect func(i, j int)
	detect = func(i, j int) {
		for d := 0; d < 4; d++ {
			x, y := direct[d][0], direct[d][1]
			if i+x >= 0 && i+x < n && j+y >= 0 && j+y < m && grid[i+x][j+y] == '1' {
				grid[i+x][j+y] = '0'
				detect(i+x, j+y)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				ans += 1
				grid[i][j] = '0'
				detect(i, j)
			}
		}
	}
	return ans
}
