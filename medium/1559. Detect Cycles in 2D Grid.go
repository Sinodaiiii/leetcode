package medium

func containsCycle260426(grid [][]byte) bool {
	directions := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	var dfs func(x, y, preDrection int) bool

	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dfs = func(x, y, preDrection int) bool {
		visited[x][y] = true
		pre := preDrection ^ 1
		for d, direction := range directions {
			if d == pre {
				continue
			}
			nx, ny := x+direction[0], y+direction[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == grid[x][y] {
				if visited[nx][ny] || dfs(nx, ny, d) {
					return true
				}
			}
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] {
				if dfs(i, j, 1) {
					return true
				}
			}
		}
	}
	return false
}
