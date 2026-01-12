package medium

import (
	"math"
	"sort"
)

func separateSquares260113(squares [][]int) float64 {
	events := make([][]int, 2*len(squares))
	total := 0
	for i, square := range squares {
		total += square[2] * square[2]
		events[2*i] = []int{square[1], square[2]}
		events[2*i+1] = []int{square[1] + square[2], -square[2]}
	}
	sort.Slice(events, func(i, j int) bool { return events[i][0] <= events[j][0] })
	target := float64(total) / 2
	currSpeed := 0
	preY := math.MinInt32
	preTotal := 0.0
	currTotal := 0.0
	// fmt.Println(events)
	for _, event := range events {
		if event[0] > preY {
			if currSpeed > 0 {
				currTotal += float64((event[0] - preY) * currSpeed)
				if currTotal > target {
					return float64(preY) + (target-preTotal)/float64(currSpeed)
				} else if currTotal == target {
					return float64(event[0])
				} else {
					preTotal = currTotal
				}
			}
			preY = event[0]
		}
		currSpeed += event[1]
		// fmt.Println(preY, preTotal, currTotal, currSpeed, target)
	}
	return 0.0
}
