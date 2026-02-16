package medium

func rob198(nums []int) int {
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

func rob198260109(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	pre, curr := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		pre, curr = curr, max(pre+nums[i], curr)
	}
	return max(pre, curr)
}

func rob198260216(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) >= 3 {
		nums[2] += nums[0]
	}
	for i := 3; i < len(nums); i++ {
		nums[i] += max(nums[i-2], nums[i-3])
	}
	return max(nums[len(nums)-1], nums[len(nums)-2])
}

func rob1982602162(nums []int) int {
	ns, s := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		ns, s = max(ns, s), ns+nums[i]
	}
	return max(ns, s)
}
