package hard

import (
	"math"
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] >= intervals[j][1]
		}
		return false
	})
	left := math.MaxInt32
	for i := len(intervals) - 1; i >= 0; i-- {
		if intervals[i][1] >= left {
			intervals = append(intervals[:i], intervals[i+1:]...)
		} else {
			left = intervals[i][1]
		}
	}
	p1, p2 := -1, -1
	ans := 0
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > p2 {
			p1, p2 = intervals[i][1]-1, intervals[i][1]
			ans += 2
		} else if intervals[i][0] <= p1 {
			continue
		} else {
			p1, p2 = p2, intervals[i][1]
			ans += 1
		}
	}
	return ans
}
