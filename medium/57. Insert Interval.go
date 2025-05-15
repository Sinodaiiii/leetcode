package medium

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}
	for i := 0; i < n; i++ {
		if newInterval[0] < intervals[i][0] {
			for j := i; j < n; j++ {
				if newInterval[1] < intervals[j][0] {
					tmp := make([][]int, n-j)
					copy(tmp, intervals[j:])
					intervals = append(intervals[:i], newInterval)
					intervals = append(intervals, tmp...)
					return intervals
				} else if newInterval[1] <= intervals[j][1] {
					tmp := make([][]int, n-j)
					copy(tmp, intervals[j:])
					intervals[j][0] = newInterval[0]
					intervals = append(intervals[:i], tmp...)
					return intervals
				}
			}
			intervals[i][0] = newInterval[0]
			intervals[i][1] = newInterval[1]
			intervals = intervals[:i+1]
			return intervals
		} else if newInterval[0] <= intervals[i][1] {
			for j := i; j < n; j++ {
				if newInterval[1] < intervals[j][0] {
					intervals[i][1] = newInterval[1]
					tmp := make([][]int, n-j)
					copy(tmp, intervals[j:])
					intervals = append(intervals[:i+1], intervals[j:]...)
					return intervals
				} else if newInterval[1] <= intervals[j][1] {
					intervals[i][1] = intervals[j][1]
					tmp := make([][]int, n-j+1)
					copy(tmp, intervals[j+1:])
					intervals = append(intervals[:i+1], intervals[j+1:]...)
					return intervals
				}
			}
			intervals[i][1] = newInterval[1]
			intervals = intervals[:i+1]
			return intervals
		} else {
			continue
		}
	}
	intervals = append(intervals, newInterval)
	return intervals
}
