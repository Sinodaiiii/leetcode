package easy

import "sort"

func minimumCost260601(cost []int) int {
	sort.Slice(cost, func(i, j int) bool { return cost[i] >= cost[j] })
	ans := 0
	for i, num := range cost {
		if i%3 != 2 {
			ans += num
		}
	}
	return ans
}
