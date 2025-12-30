package medium

func numMagicSquaresInside(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	pair := map[int]int{4: 6, 3: 7, 8: 2, 1: 9, 6: 4, 7: 3, 2: 8, 9: 1}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 5 {
				if grid[i-1][j-1] <= 0 || grid[i-1][j-1] >= 10 || grid[i-1][j] <= 0 || grid[i-1][j] >= 10 || grid[i-1][j+1] <= 0 || grid[i-1][j+1] >= 10 || grid[i][j+1] <= 0 || grid[i][j+1] >= 10 {
					continue
				}
				if grid[i-1][j-1] == pair[grid[i+1][j+1]] && grid[i-1][j] == pair[grid[i+1][j]] && grid[i-1][j+1] == pair[grid[i+1][j-1]] && grid[i][j+1] == pair[grid[i][j-1]] {
					if grid[i-1][j-1]+grid[i-1][j]+grid[i-1][j+1] == grid[i][j-1]+grid[i][j]+grid[i][j+1] && grid[i][j-1]+grid[i][j]+grid[i][j+1] == grid[i+1][j-1]+grid[i+1][j]+grid[i+1][j+1] {
						if grid[i-1][j-1]+grid[i][j-1]+grid[i+1][j-1] == grid[i-1][j]+grid[i][j]+grid[i+1][j] && grid[i-1][j]+grid[i][j]+grid[i+1][j] == grid[i-1][j+1]+grid[i][j+1]+grid[i+1][j+1] {
							ans += 1
						}
					}
				}
			}
		}
	}
	return ans
}
