package hot100

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	ans := []int{}
	stack := [][2]int{}
	for i := 0; i < n; i++ {
		if len(stack) > 0 && stack[0][0] <= i-k {
			stack = stack[1:]
		}
		for len(stack) > 0 && nums[i] >= stack[len(stack)-1][1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, [2]int{i, nums[i]})
		if i >= k-1 {
			ans = append(ans, stack[0][1])
		}
	}
	return ans
}
