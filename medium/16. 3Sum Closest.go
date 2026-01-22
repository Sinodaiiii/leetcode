package medium

import (
	"math"
	"sort"
)

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func threeSumClosest260123(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	gap := math.MaxInt32
	ans := target
	for i := 1; i < n-1; i++ {
		l, r := i-1, i+1
		for l >= 0 && r < n {
			curr := nums[l] + nums[i] + nums[r]
			if abs(curr-target) < gap {
				gap = abs(curr - target)
				ans = curr
			}
			if curr < target {
				r += 1
			} else {
				l -= 1
			}
		}
	}
	return ans
}
