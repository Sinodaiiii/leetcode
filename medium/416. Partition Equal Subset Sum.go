package medium

func canPartition(nums []int) bool {
	sum, n := 0, len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	dp := make([]int, sum+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := sum; j >= 0; j-- {
			if dp[j] == 1 && j+nums[i] <= sum {
				dp[j+nums[i]] = 1
			}
		}
		if dp[sum] == 1 {
			return true
		}
	}
	return false
}
