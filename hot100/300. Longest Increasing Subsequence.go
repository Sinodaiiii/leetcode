package hot100

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := []int{nums[0]}
	var findNum func(target int) int
	findNum = func(target int) int {
		l, r := 0, len(dp)-1
		mid := (l + r) / 2
		for l < r {
			if dp[mid] < target {
				l = mid + 1
			} else if dp[mid] > target {
				if mid == 0 || dp[mid-1] < target {
					return mid
				}
				r = mid - 1
			} else {
				return mid
			}
			mid = (l + r) / 2
		}
		return l
	}

	for i := 1; i < n; i++ {
		if nums[i] > dp[len(dp)-1] {
			dp = append(dp, nums[i])
		} else {
			index := findNum(nums[i])
			if nums[i] < dp[index] {
				dp[index] = nums[i]
			}
		}
	}
	return len(dp)
}
