package hot100

func findTargetSumWays(nums []int, target int) int {
	if target < 0 {
		target = -target
	}
	total := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		total += nums[i]
	}
	if (total+target)%2 == 1 {
		return 0
	}
	target = (total + target) / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := target; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] = dp[j] + dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
