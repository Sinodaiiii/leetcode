package medium

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] <= intervals[j][0] })
	ans := make([][]int, 0)
	for _, value := range intervals {
		if len(ans) == 0 {
			ans = append(ans, value)
		}
		r := ans[len(ans)-1][1]
		if r < value[0] {
			ans = append(ans, value)
		} else if r >= value[0] {
			ans[len(ans)-1][1] = max(r, value[1])
		}
	}
	return ans
}
