package medium

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans := n + 1
	l, r := 0, -1
	sum := 0
	for r != n-1 {
		r += 1
		sum += nums[r]
		for sum >= target {
			ans = min(ans, r-l+1)
			sum -= nums[l]
			l += 1
		}
	}
	if ans == n+1 {
		return 0
	}
	return ans
}
