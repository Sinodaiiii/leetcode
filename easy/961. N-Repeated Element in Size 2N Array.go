package easy

import "math"

func repeatedNTimes260102(nums []int) int {
	if nums[0] == nums[len(nums)-1] {
		return nums[0]
	}
	pre := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] == pre {
			return pre
		}
		if i < len(nums)-2 && nums[i] == nums[i+2] {
			return nums[i]
		}
		pre = nums[i]
	}
	return 0
}
