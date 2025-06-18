package hot100

func moveZeroes(nums []int) {
	n := len(nums)
	curr := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[curr] = nums[i]
			curr += 1
		}
	}
	for ; curr < n; curr++ {
		nums[curr] = 0
	}
}
