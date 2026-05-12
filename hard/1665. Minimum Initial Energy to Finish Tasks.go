package hard

import "sort"

func minimumEffort260512(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i][1]-tasks[i][0] > tasks[j][1]-tasks[j][0] {
			return true
		} else if tasks[i][1]-tasks[i][0] == tasks[j][1]-tasks[j][0] {
			return tasks[i][1] > tasks[j][1]
		}
		return false
	})
	total, maxTime := 0, 0
	for _, task := range tasks {
		maxTime = max(maxTime, task[1])
		total += task[0]
	}
	l, r := maxTime-total, total+maxTime
	check := func(x int) bool {
		for _, task := range tasks {
			if x < task[1] {
				return false
			}
			x -= task[0]
		}
		return true
	}
	target := r
	for l <= r {
		mid := (l + r) / 2
		if check(mid) {
			target = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return target
}
