package hot100

func longestConsecutive(nums []int) int {
	dict := map[int]bool{}
	n := len(nums)
	for i := 0; i < n; i++ {
		dict[nums[i]] = true
	}
	ans, curr := 0, 0
	for i := 0; i < n; i++ {
		if dict[nums[i]] == true {
			curr = 1
			delete(dict, nums[i])
			for j := nums[i] - 1; ; j-- {
				if dict[j] == true {
					curr += 1
					delete(dict, j)
				} else {
					break
				}
			}
			for j := nums[i] + 1; ; j++ {
				if dict[j] == true {
					curr += 1
					delete(dict, j)
				} else {
					break
				}
			}
			ans = max(ans, curr)
		}
	}
	return ans
}
