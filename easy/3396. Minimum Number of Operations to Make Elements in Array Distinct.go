package easy

func minimumOperations(nums []int) int {
	n := len(nums)
	dict := map[int]bool{}
	index := -1
	for i := n - 1; i >= 0; i-- {
		if dict[nums[i]] == true {
			index = i
			break
		}
		dict[nums[i]] = true
	}
	if index < 0 {
		return 0
	}
	return (index + 3) / 3
}
