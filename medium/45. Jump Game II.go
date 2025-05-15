package medium

import "math"

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	dp := make([]int, 1001)
	for i := 1; i < 1001; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i < n; i++ {
		if i < 1001 {
			for j := 0; j < i; j++ {
				if j+nums[j] >= i {
					dp[i] = min(dp[j]+1, dp[i])
				}
			}
		} else {
			dp[i%1001] = math.MaxInt32
			for j := i + 1; j < i+1001; j++ {
				if (j - 1001 + nums[j-1001]) >= i {
					dp[i%1001] = min(dp[j%1001]+1, dp[i%1001])
				}
			}
		}
	}
	return dp[(n-1)%1001]
}
