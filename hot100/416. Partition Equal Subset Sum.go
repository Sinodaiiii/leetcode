package hot100

func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	dp := make([]bool, sum/2+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := sum / 2; j >= nums[i]; j-- {
			dp[j] = dp[j-nums[i]] || dp[j]
		}
	}
	return dp[sum/2]
}
