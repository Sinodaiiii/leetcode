package medium

import "sort"

func maximizeSquareArea260116(m int, n int, hFences []int, vFences []int) int {
	hFences = append(hFences, 1, m)
	vFences = append(vFences, 1, n)
	sort.Ints(hFences)
	sort.Ints(vFences)
	dict := map[int]bool{}
	for i := 0; i < len(hFences); i++ {
		for j := i + 1; j < len(hFences); j++ {
			dict[hFences[j]-hFences[i]] = true
		}
	}
	// fmt.Println(dict)
	currMax := 0
	for i := 0; i < len(vFences); i++ {
		for j := len(vFences) - 1; j > i && vFences[j]-vFences[i] > currMax; j-- {
			if dict[vFences[j]-vFences[i]] {
				currMax = vFences[j] - vFences[i]
			}
		}
	}
	if currMax == 0 {
		return -1
	}
	return currMax * currMax % 1000000007
}
