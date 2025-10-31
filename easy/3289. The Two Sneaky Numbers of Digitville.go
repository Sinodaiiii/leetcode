package easy

func getSneakyNumbers(nums []int) []int {
	m := len(nums)
	n := m - 2
	ans := make([]int, 0, 2)
	for _, x := range nums {
		idx := x % n
		if nums[idx] >= n {
			ans = append(ans, x%n)
		} else {
			nums[idx] += n
		}
	}
	return ans
}
