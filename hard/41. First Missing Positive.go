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

func firstMissingPositive251229(nums []int) int {
	nums = append(nums, 0, 0)
	n := len(nums)
	for i, num := range nums {
		if num <= 0 {
			nums[i] -= n
		} else {
			nums[i] = -nums[i]
		}
	}
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if num < n && nums[num] < 0 {
			nums[num] = -nums[num]
		}
	}
	for i := 1; i < n; i++ {
		if nums[i] <= 0 {
			return i
		}
	}
	return 0
}
