package medium

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	ans := nums[0]
	pre := nums[0]
	for i := 1; i < n; i++ {
		pre = max(pre+nums[i], nums[i])
		ans = max(pre, ans)
	}

	l, r := make([]int, n), make([]int, n)
	l[0] = nums[0]
	r[n-1] = nums[n-1]
	pre = nums[0]
	for i := 1; i < n; i++ {
		pre += nums[i]
		l[i] = max(pre, l[i-1])
	}
	pre = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		pre += nums[i]
		r[i] = max(pre, r[i+1])
	}
	for i := 0; i < n-1; i++ {
		ans = max(ans, l[i]+r[i+1])
	}
	return ans
}
