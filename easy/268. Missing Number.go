package easy

func missingNumber(nums []int) int {
	n := len(nums)
	nums = append(nums, -1)
	for i := 0; i < n; i++ {
		if nums[i] == -2 {
			continue
		} else {
			next := nums[i]
			for next >= 0 {
				next, nums[next] = nums[next], -1
			}
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] >= 0 {
			return i
		}
	}
	return n
}
