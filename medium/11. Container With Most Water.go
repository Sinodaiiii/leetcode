package medium

import "math"

func maxArea(height []int) int {
	n := len(height)
	l, r := 0, n-1
	ans := math.MinInt32
	for l < r {
		ans = max(ans, (r-l)*min(height[l], height[r]))
		if height[l] < height[r] {
			l += 1
		} else if height[l] > height[r] {
			r -= 1
		} else {
			if height[l+1] > height[r-1] {
				l += 1
			} else {
				r -= 1
			}
		}
	}
	return ans
}
