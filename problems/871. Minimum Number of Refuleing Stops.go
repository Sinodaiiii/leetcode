package problems

import (
	"container/heap"
	"sort"
)

func minRefuelStops(target, startFuel int, stations [][]int) int {
	n := len(stations)
	dp := make([]int, n+1)
	dp[0] = startFuel
	for i, s := range stations {
		for j := i; j >= 0; j-- {
			if dp[j] >= s[0] {
				dp[j+1] = max871(dp[j+1], dp[j]+s[1])
			}
		}
	}
	for i, v := range dp {
		if v >= target {
			return i
		}
	}
	return -1
}

func max871(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func minRefuelStops2(target, startFuel int, stations [][]int) (ans int) {
	fuel, prev, h := startFuel, 0, hp{}
	for i, n := 0, len(stations); i <= n; i++ {
		curr := target
		if i < n {
			curr = stations[i][0]
		}
		fuel -= curr - prev
		for fuel < 0 && h.Len() > 0 {
			fuel += heap.Pop(&h).(int)
			ans++
		}
		if fuel < 0 {
			return -1
		}
		if i < n {
			heap.Push(&h, stations[i][1])
			prev = curr
		}
	}
	return
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
