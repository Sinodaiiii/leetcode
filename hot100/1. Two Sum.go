package hot100

func twoSum(nums []int, target int) []int {
	dict := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if index, exist := dict[target-nums[i]]; exist {
			return []int{index, i}
		}
		dict[nums[i]] = i
	}
	return []int{}
}
