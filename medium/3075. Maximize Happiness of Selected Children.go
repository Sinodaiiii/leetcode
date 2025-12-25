package medium

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool { return happiness[i] >= happiness[j] })
	ans := 0
	for i := 0; i < k; i++ {
		val := happiness[i] - i
		if val <= 0 {
			break
		}
		ans += val
	}
	return int64(ans)
}
