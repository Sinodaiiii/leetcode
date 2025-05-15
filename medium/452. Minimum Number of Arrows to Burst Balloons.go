package medium

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] <= points[j][1]
	})

	ans := 1
	curr := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > curr {
			curr = points[i][1]
			ans++
		}
	}
	return ans
}
