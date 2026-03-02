package medium

func minSwaps260302(grid [][]int) int {
	n := len(grid)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		curr := 0
		for j := n - 1; j >= 0 && grid[i][j] == 0; j-- {
			curr += 1
		}
		nums[i] = curr
	}
	used := make([]bool, n)
	ans := 0
	for i := 0; i < n-1; i++ {
		curr := 0
		found := false
		for j := 0; j < n; j++ {
			if used[j] {
				continue
			}
			if nums[j] >= n-1-i {
				used[j] = true
				ans += curr
				found = true
				break
			}
			curr += 1
		}
		if !found {
			return -1
		}
	}
	return ans
}
