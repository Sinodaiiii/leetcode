package hot100

import "sort"

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] <= intervals[j][0] })
	n := len(intervals)
	list := make([]int, 0)
	ans := 0
	for i := 0; i < n; i++ {
		curr := intervals[i][0]
		for j := len(list) - 1; j >= 0; j-- {
			if list[j] <= curr {
				list = append(list[:j], list[j+1:]...)
			}
		}
		list = append(list, intervals[i][1])
		ans = max(ans, len(list))
	}
	return ans
}
