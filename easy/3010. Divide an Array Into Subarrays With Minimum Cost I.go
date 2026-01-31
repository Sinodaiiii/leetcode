package easy

import "math"

func minimumCost260201(nums []int) int {
	a, b := math.MaxInt32, math.MaxInt32
	for i, num := range nums {
		if i == 0 {
			continue
		}
		if num < max(a, b) {
			a, b = min(a, b), num
		}
	}
	return nums[0] + a + b
}
