package hot100

func merge(intervals [][]int) [][]int {
	var sortIntervals func(l, r int)
	sortIntervals = func(l, r int) {
		if l >= r {
			return
		}
		pivot := r
		curr := l
		for i := l; i < r; i++ {
			if intervals[i][0] < intervals[pivot][0] {
				intervals[curr], intervals[i] = intervals[i], intervals[curr]
				curr += 1
			}
		}
		intervals[curr], intervals[pivot] = intervals[pivot], intervals[curr]
		sortIntervals(l, curr-1)
		sortIntervals(curr+1, r)
	}
	// sort.Slice(intervals, func (i, j int) bool {return intervals[i][0] <= intervals[j][0]})
	n := len(intervals)
	sortIntervals(0, n-1)
	curr := 0
	for i := 1; i < n; i++ {
		if intervals[i][0] <= intervals[curr][1] {
			intervals[curr][1] = max(intervals[curr][1], intervals[i][1])
		} else {
			curr += 1
			intervals[curr][0], intervals[curr][1] = intervals[i][0], intervals[i][1]
		}
	}
	return intervals[:curr+1]
}
