package easy

import "math"

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	ml := make([]int, n)
	ans := 0
	for i := 1; i < n; i++ {
		ml[i] = max(ml[i-1], nums[i-1])
	}
	mr := math.MinInt64
	for i := n - 2; i >= 1; i-- {
		mr = max(mr, nums[i+1])
		ans = max(ans, mr*(ml[i]-nums[i]))
	}
	return int64(ans)
}
