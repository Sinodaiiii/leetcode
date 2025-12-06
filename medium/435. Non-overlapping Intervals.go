package medium

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans := 0
	pre := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[pre][1] {
			pre = i
		} else {
			ans += 1
			if intervals[i][1] < intervals[pre][1] {
				pre = i
			}
		}
	}
	return ans
}
