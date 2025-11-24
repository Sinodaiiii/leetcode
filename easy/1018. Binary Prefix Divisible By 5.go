package easy

func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))
	ans[0] = (nums[0] == 0)
	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		curr = curr<<1 + nums[i]
		curr = curr % 10
		ans[i] = (curr%5 == 0)
	}
	return ans
}
