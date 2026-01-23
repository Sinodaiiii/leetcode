package medium

import "sort"

func minPairSum260124(nums []int) int {
	sort.Ints(nums)
	ans := 0
	n := len(nums)
	for i := 0; i < n/2; i++ {
		ans = max(ans, nums[i]+nums[n-i-1])
	}
	return ans
}
