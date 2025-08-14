package medium

func findDuplicates(nums []int) []int {
	ans := []int{}
	n := len(nums)
	nums = append(nums, 0)
	for i := 0; i < n; i++ {
		if nums[i] == -1 {
			continue
		}
		curr := nums[i]
		nums[i] = -2
		for curr > 0 {
			if nums[curr] == -1 {
				ans = append(ans, curr)
				break
			}
			curr, nums[curr] = nums[curr], -1
		}
	}
	return ans
}
