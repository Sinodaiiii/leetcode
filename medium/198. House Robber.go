package medium

func rob(nums []int) int {
	n := len(nums)
	yes := make([]int, n)
	no := make([]int, n)
	yes[0] = nums[0]
	for i := 1; i < n; i++ {
		yes[i] = max(yes[i-1], no[i-1]+nums[i])
		no[i] = max(yes[i-1], no[i-1])
	}
	return max(yes[n-1], no[n-1])
}
