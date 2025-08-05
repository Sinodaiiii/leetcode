package medium

func orangesRotting(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])
	distance := make([][]int, m)
	for i := 0; i < m; i++ {
		distance[i] = make([]int, n)
	}
	direction := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var rot func(x, y int)
	rot = func(x, y int) {
		if grid[x][y] == 0 {
			return
		} else if grid[x][y] == 1 {
			for _, dir := range direction {
				i, j := x+dir[0], y+dir[1]
				if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == 1 {
					if distance[i][j] == 0 {
						distance[i][j] = distance[x][y] + 1
						rot(i, j)
					} else if distance[i][j] > distance[x][y]+1 {
						distance[i][j] = distance[x][y] + 1
						rot(i, j)
					}
				}
			}
		} else if grid[x][y] == 2 {
			for _, dir := range direction {
				i, j := x+dir[0], y+dir[1]
				if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == 1 {
					if distance[i][j] != 1 {
						distance[i][j] = 1
						rot(i, j)
					}
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				rot(i, j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && distance[i][j] == 0 {
				return -1
			}
			ans = max(ans, distance[i][j])
		}
	}
	return ans
}
