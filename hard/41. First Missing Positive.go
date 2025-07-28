package hard

import "math"

func firstMissingPositive(nums []int) int {
	nums = append(nums, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] >= 0 && nums[i] < n {
			curr := nums[i]
			for curr >= 0 && curr < n {
				nums[curr], curr = math.MinInt32-100, nums[curr]
			}
		}
	}
	for i := 1; i < n; i++ {
		if nums[i] >= math.MinInt32 {
			return i
		}
	}
	return n
}
