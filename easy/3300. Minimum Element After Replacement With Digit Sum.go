package easy

import "math"

func minElement260529(nums []int) int {
	ans := math.MaxInt32
	for _, num := range nums {
		curr := 0
		for num != 0 {
			curr += num % 10
			num = num / 10
		}
		ans = min(ans, curr)
	}
	return ans
}
