package medium

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum < target {
		return 0
	}
	l, r := 0, n-1
	for {
		if sum == target {
			return r - l + 1
		}
		if nums[l] <= nums[r] {
			sum -= nums[l]
			if sum < target {
				return r - l + 1
			}
			l++
		} else {
			sum -= nums[r]
			if sum < target {
				return r - l + 1
			}
			r--
		}
	}
}
