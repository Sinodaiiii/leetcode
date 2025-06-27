package hot100

func canJump(nums []int) bool {
	n := len(nums) - 1
	curr, r := 0, 0
	for curr <= r {
		if curr+nums[curr] >= r {
			r = curr + nums[curr]
			if r >= n {
				return true
			}
		}
		curr++
	}
	return false
}
