package easy

func constructTransformedArray260205(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i, num := range nums {
		if num < 0 {
			num = num%n + n
		}
		ans[i] = nums[(i+num)%n]
	}
	return ans
}
