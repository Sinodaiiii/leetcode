package hot100

import "math"

func majorityElement(nums []int) int {
	curr := math.MaxInt64
	count := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			curr = nums[i]
			count += 1
		} else if nums[i] == curr {
			count += 1
		} else {
			count -= 1
		}
	}
	return curr
}
