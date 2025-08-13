package medium

import "sort"

func deleteAndEarn(nums []int) int {
	dp := [2]int{}
	sort.Slice(nums, func(i, j int) bool { return nums[i] <= nums[j] })
	list := make([][2]int, 0)
	for i := 0; i < len(nums); i++ {
		if len(list) == 0 || nums[i] != list[len(list)-1][0] {
			list = append(list, [2]int{nums[i], 1})
		} else {
			list[len(list)-1][1] += 1
		}
	}

	for i := 0; i < len(list); i++ {
		if i > 0 && list[i][0] == list[i-1][0]+1 {
			dp[0], dp[1] = dp[1]+list[i][0]*list[i][1], max(dp[0], dp[1])
		} else {
			dp[0], dp[1] = max(dp[0], dp[1])+list[i][0]*list[i][1], max(dp[0], dp[1])
		}
	}

	return max(dp[0], dp[1])
}
