package medium

func maxSubArray(nums []int) int {
	n := len(nums)
	dp, ans := nums[0], nums[0]
	for i := 1; i < n; i++ {
		dp = max(dp+nums[i], nums[i])
		ans = max(ans, dp)
	}
	return ans
}

func maxSubArray2(nums []int) int {
	sum := 0
	m := 0
	ans := nums[0]
	for _, num := range nums {
		sum += num
		ans = max(ans, sum-m)
		m = min(m, sum)
	}
	return ans
}
