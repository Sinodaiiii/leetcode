package medium

import "sort"

func minOperations260428(grid [][]int, x int) int {
	m, n := len(grid), len(grid[0])
	total := m * n
	line := make([]int, 0, total)
	pivot := grid[0][0]
	for i := range grid {
		for j := range grid[0] {
			if (grid[i][j]-pivot)%x != 0 {
				return -1
			}
			line = append(line, grid[i][j])
		}
	}
	sort.Ints(line)
	ans := 0
	for i := 1; i < total; i++ {
		ans += line[i] - line[0]
	}
	curr, right := ans, total-1
	for i := 1; i < total; i++ {
		curr = curr + (total-2*right)*(line[i]-line[i-1])
		ans = min(ans, curr)
		right -= 1
	}
	return ans / x
}
