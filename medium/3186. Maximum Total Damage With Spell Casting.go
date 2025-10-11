package medium

import "sort"

func maximumTotalDamage(power []int) int64 {
	dict := map[int]int{}
	keys := []int{}
	for i := 0; i < len(power); i++ {
		if dict[power[i]] == 0 {
			keys = append(keys, power[i])
		}
		dict[power[i]] += power[i]
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] <= keys[j] })
	dp := [3]int{}
	pre := -1
	for i := 0; i < len(keys); i++ {
		// fmt.Println(dp)
		if keys[i]-pre == 1 {
			dp[0], dp[1], dp[2] = max(dp[0], dp[1]), dp[2], max(dp[2], dp[0]+dict[keys[i]])
		} else if keys[i]-pre == 2 {
			dp[0], dp[1], dp[2] = max(max(dp[1], dp[2]), dp[0]), dp[2], max(max(dp[2], dp[0]+dict[keys[i]]), dp[1]+dict[keys[i]])
		} else {
			dp[0], dp[1], dp[2] = max(max(dp[1], dp[2]), dp[0]), dp[2], max(max(dp[1], dp[2]), dp[0])+dict[keys[i]]
		}
		pre = keys[i]
	}
	// fmt.Println(dp)
	return int64(max(max(dp[1], dp[2]), dp[0]))
}
