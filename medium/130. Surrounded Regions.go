package medium

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	box := make([][]bool, m)
	for i := 0; i < m; i++ {
		box[i] = make([]bool, n)
	}
	direction := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		box[i][j] = true
		for _, d := range direction {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && y >= 0 && y < n {
				if board[x][y] == 'O' && box[x][y] == false {
					dfs(x, y)
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if board[0][i] == 'O' && box[0][i] == false {
			dfs(0, i)
		}
		if board[m-1][i] == 'O' && box[m-1][i] == false {
			dfs(m-1, i)
		}
	}
	for i := 1; i < m; i++ {
		if board[i][0] == 'O' && box[i][0] == false {
			dfs(i, 0)
		}
		if board[i][n-1] == 'O' && box[i][n-1] == false {
			dfs(i, n-1)
		}
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' && box[i][j] == false {
				board[i][j] = 'X'
			}
		}
	}
}
