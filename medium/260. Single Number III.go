package medium

func singleNumber(nums []int) []int {
	if len(nums) == 2 {
		return nums
	}
	dict := map[int]bool{}
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		if !dict[nums[i]] {
			dict[nums[i]] = true
		} else {
			delete(dict, nums[i])
		}
	}
	for key, _ := range dict {
		ans = append(ans, key)
	}
	return ans
}
