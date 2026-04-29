package easy

import "math"

func getMinDistance260429(nums []int, target int, start int) int {
	ans := math.MaxInt32
	abs := func(x int) int {
		if x > 0 {
			return x
		}
		return -x
	}
	for i, num := range nums {
		if num == target {
			ans = min(ans, abs(start-i))
		}
	}
	return ans
}
