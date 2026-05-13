package medium

import "sort"

func minMoves260513(nums []int, limit int) int {
	n := len(nums)
	change := make([][2]int, 0, n*2)
	for i := 0; i < n/2; i++ {
		change = append(change, [2]int{min(nums[i], nums[n-1-i]) + 1, -1})
		change = append(change, [2]int{nums[i] + nums[n-1-i], -1})
		change = append(change, [2]int{nums[i] + nums[n-1-i] + 1, +1})
		change = append(change, [2]int{max(nums[i], nums[n-1-i]) + limit + 1, +1})
	}
	ans, curr := n, n
	sort.Slice(change, func(i, j int) bool { return change[i][0] <= change[j][0] })
	cn := len(change)
	for i := 0; i < cn; {
		pre := change[i][0]
		for i < cn && change[i][0] == pre {
			curr += change[i][1]
			i += 1
		}
		ans = min(ans, curr)
	}
	return ans
}
