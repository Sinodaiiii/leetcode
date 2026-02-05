package medium

import (
	"math"
	"sort"
)

func minRemoval260206(nums []int, k int) int {
	sort.Ints(nums)
	l, r := 0, 0
	n := len(nums)
	ans := math.MaxInt32
	for r < n && l < n {
		for r < n && nums[r] <= nums[l]*k {
			r += 1
		}
		ans = min(ans, n-(r-l))
		l += 1
	}
	return ans
}
