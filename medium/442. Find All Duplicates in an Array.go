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

func findDuplicates260206(nums []int) []int {
	ans := []int{}
	for i, num := range nums {
		if num <= 0 {
			continue
		}
		curr := num
		nums[i] = 0
		for curr > 0 {
			if nums[curr-1] < 0 {
				ans = append(ans, curr)
			}
			curr, nums[curr-1] = nums[curr-1], -1
		}
	}
	return ans
}
