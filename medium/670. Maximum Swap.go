package medium

import (
	"sort"
)

func maximumSwap(num int) int {
	if num <= 9 {
		return num
	}
	origin := make([]int, 0)
	for num != 0 {
		origin = append([]int{num % 10}, origin...)
		num = num / 10
	}
	n := len(origin)
	t := make([]int, n)
	copy(t, origin)
	sort.Slice(t, func(i, j int) bool {
		return t[i] > t[j]
	})
	for i := 0; i < n; i++ {
		if t[i] != origin[i] {
			tmp := origin[i]
			origin[i] = t[i]
			for j := n - 1; j >= 0; j-- {
				if origin[j] == t[i] {
					origin[j] = tmp
					break
				}
			}
			break
		}
	}
	ret := 0
	for i := 0; i < n; i++ {
		ret *= 10
		ret += origin[i]
	}
	return ret
}
