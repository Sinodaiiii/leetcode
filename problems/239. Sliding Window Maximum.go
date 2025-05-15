package problems

import "math"

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 1 {
		return []int{1}
	} else if n <= k {
		ans := math.MinInt32
		for i := 0; i < n; i++ {
			ans = max(ans, nums[i])
		}
		return []int{ans}
	}
	stack := make([]int, n)
	stack[0] = 0
	for i := 1; i < n; i++ {
		for j := 1; j < k; j++ {
			if i-j >= 0 && nums[i-j] > nums[i] {
				stack[i] = i - j
			}
		}
	}
	ans := make([]int, 0)
	for i := k - 1; i < n; i++ {
		j := i
		index := stack[i]
		for index != 0 {
			j = i - stack[i]
			index = stack[j]
		}
		ans = append(ans, nums[j])
	}
	return ans
}
