package medium

import "sort"

func maxDistinctElements(nums []int, k int) int {
	// sort.Slice(nums, func(i, j int) bool { return nums[i] <= nums[j]})
	sort.Ints(nums)
	left := nums[0]
	ans := 0
	for _, num := range nums {
		left = max(left, num)
		if left <= num+2*k {
			left += 1
			ans += 1
		}
	}
	return ans
}
