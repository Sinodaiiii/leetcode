package medium

func exist(board [][]byte, word string) bool {
	step := map[int][2]int{
		0: {-1, 0},
		1: {0, 1},
		2: {1, 0},
		3: {0, -1},
	}
	nw := len(word)
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var search func(x, y int, length int) bool
	search = func(x, y int, length int) bool {
		if length == nw {
			return true
		}
		for c := 0; c < 4; c++ {
			xs, ys := step[c][0], step[c][1]
			if x+xs >= 0 && x+xs < m && y+ys >= 0 && y+ys < n {
				if !visited[x+xs][y+ys] && board[x+xs][y+ys] == word[length] {
					visited[x+xs][y+ys] = true
					if search(x+xs, y+ys, length+1) {
						return true
					}
					visited[x+xs][y+ys] = false
				}
			}
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				visited[i][j] = true
				if search(i, j, 1) {
					return true
				}
				visited[i][j] = false
			}
		}
	}
	return false
}
