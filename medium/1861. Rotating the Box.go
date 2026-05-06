package medium

func rotateTheBox(boxGrid [][]byte) [][]byte {
	m, n := len(boxGrid), len(boxGrid[0])
	ans := make([][]byte, n)
	for i := range ans {
		ans[i] = make([]byte, m)
	}
	for i := 0; i < m; i++ {
		next := n - 1
		for j := n - 1; j >= 0; j-- {
			switch boxGrid[i][j] {
			case '.':
				continue
			case '#':
				ans[next][m-i-1] = '#'
				next -= 1
			case '*':
				for ; next > j; next-- {
					ans[next][m-i-1] = '.'
				}
				ans[next][m-i-1] = '*'
				next -= 1
			}
		}
		for ; next >= 0; next-- {
			ans[next][m-i-1] = '.'
		}
	}
	return ans
}
