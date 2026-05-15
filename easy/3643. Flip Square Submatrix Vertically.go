package easy

func reverseSubmatrix260515(grid [][]int, x int, y int, k int) [][]int {
	for i := 0; i < k/2; i++ {
		for j := y; j < y+k; j++ {
			grid[x+i][j], grid[x+k-1-i][j] = grid[x+k-1-i][j], grid[x+i][j]
		}
	}
	return grid
}
