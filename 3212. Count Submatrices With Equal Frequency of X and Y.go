package main

func numberOfSubmatrices260319(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	num := make([][2]int, n)
	ans := 0
	for i := 0; i < m; i++ {
		pre := [2]int{0, 0}
		for j := 0; j < n; j++ {
			if grid[i][j] == 'X' {
				num[j][0] += 1
			} else if grid[i][j] == 'Y' {
				num[j][1] += 1
			}
			pre[0] += num[j][0]
			pre[1] += num[j][1]
			if pre[0] > 0 && pre[0] == pre[1] {
				ans += 1
			}
		}
	}
	return ans
}
