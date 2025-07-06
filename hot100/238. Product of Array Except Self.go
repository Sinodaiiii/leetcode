package hot100

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[0] = 1
	left := 1
	for i := 1; i < n; i++ {
		left *= nums[i-1]
		ans[i] = left
	}
	right := 1
	for i := n - 2; i >= 0; i-- {
		right *= nums[i+1]
		ans[i] *= right
	}
	return ans
}
