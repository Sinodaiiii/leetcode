package problems

func solveNQueens(n int) [][]string {
	ans := [][]string{}
	curr := make([][]byte, n)
	for i := 0; i < n; i++ {
		curr[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			curr[i][j] = '.'
		}
	}
	line := make([]bool, n)
	var dfs func(level int)
	dfs = func(level int) {
		if level == n {
			tmp := make([]string, n)
			for i := 0; i < n; i++ {
				tmp[i] = string(curr[i])
			}
			ans = append(ans, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if line[i] == true {
				continue
			}
			flag := 0
			for j := 1; j <= level; j++ {
				if level-j >= 0 && i-j >= 0 && curr[level-j][i-j] == 'Q' {
					flag = 1
					break
				}
				if level-j >= 0 && i+j < n && curr[level-j][i+j] == 'Q' {
					flag = 1
					break
				}
			}
			if flag == 1 {
				continue
			}
			curr[level][i] = 'Q'
			line[i] = true
			dfs(level + 1)
			curr[level][i] = '.'
			line[i] = false
		}
	}

	dfs(0)
	return ans
}
