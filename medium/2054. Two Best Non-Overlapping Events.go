package medium

import "sort"

func maxTwoEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool { return events[i][0] <= events[j][0] })
	n := len(events)
	right := [][2]int{[2]int{events[n-1][0], events[n-1][2]}}
	for i := n - 2; i >= 0; i-- {
		if events[i][0] == right[len(right)-1][0] {
			right[len(right)-1][1] = max(right[len(right)-1][1], events[i][2])
		} else {
			right = append(right, [2]int{events[i][0], max(events[i][2], right[len(right)-1][1])})
		}
	}
	sort.Slice(events, func(i, j int) bool { return events[i][1] <= events[j][1] })
	ans := 0
	rIndex := len(right) - 1
	for _, event := range events {
		for rIndex >= 0 && right[rIndex][0] <= event[1] {
			rIndex -= 1
		}
		rValue := 0
		if rIndex >= 0 {
			rValue = right[rIndex][1]
		}
		ans = max(event[2]+rValue, ans)
	}
	return ans
}
