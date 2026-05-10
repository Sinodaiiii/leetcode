package medium

func maximumJumps260510(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n-1; i++ {
		dp[i] = -1
	}
	for i := n - 2; i >= 0; i-- {
		hop := -1
		for j := i + 1; j < n; j++ {
			if abs(nums[i]-nums[j]) <= target {
				hop = max(hop, dp[j])
			}
		}
		if hop != -1 {
			dp[i] = hop + 1
		}
	}

	return dp[0]
}
