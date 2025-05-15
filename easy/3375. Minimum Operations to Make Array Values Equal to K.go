package easy

func minOperations(nums []int, k int) int {
	dict := map[int]bool{}
	dict[k] = true
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] < k {
			return -1
		}
		dict[nums[i]] = true
	}
	return len(dict) - 1
}
