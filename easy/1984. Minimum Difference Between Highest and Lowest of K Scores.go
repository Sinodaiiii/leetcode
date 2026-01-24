package easy

import (
	"math"
	"sort"
)

func minimumDifference260125(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	sort.Ints(nums)
	ans := math.MaxInt32
	for i := 0; i <= len(nums)-k; i++ {
		ans = min(ans, nums[i+k-1]-nums[i])
	}
	return ans
}
