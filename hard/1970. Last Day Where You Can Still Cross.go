package hard

func latestDayToCross(row int, col int, cells [][]int) int {
	matrix := make([][]int, row)
	for i := 0; i < row; i++ {
		matrix[i] = make([]int, col)
	}
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == row-1 {
			return true
		}
		visited[i][j] = true
		ret := false
		for _, direction := range directions {
			x, y := i+direction[0], j+direction[1]
			if x >= 0 && x < row && y >= 0 && y < col && matrix[x][y] == 0 && !visited[x][y] {
				ret = dfs(x, y)
				if ret {
					ret = true
					break
				}
			}
		}
		return ret
	}

	check := func() bool {
		for i := 0; i < row; i++ {
			visited[i] = make([]bool, col)
		}
		for i := 0; i < col; i++ {
			if matrix[0][i] == 0 && dfs(0, i) {
				return true
			}
		}
		return false
	}

	target := 0
	direct, pre := 1, 0
	l, r := 0, len(cells)-1
	for l <= r {
		mid := (l + r) / 2
		if direct == 1 {
			for i := pre; i <= mid; i++ {
				x, y := cells[i][0]-1, cells[i][1]-1
				matrix[x][y] = 1
			}
		} else {
			for i := pre; i > mid; i-- {
				x, y := cells[i][0]-1, cells[i][1]-1
				matrix[x][y] = 0
			}
		}
		if check() {
			target = mid
			l = mid + 1
			direct = 1
		} else {
			r = mid - 1
			direct = -1
		}
		pre = mid
	}
	return target + 1
}
