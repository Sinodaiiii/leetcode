package easy

func isGood260514(nums []int) bool {
	n := len(nums)
	dict := make([]int, n)
	n -= 1
	for _, num := range nums {
		if num > n || (num == n && dict[num] == 2) || (num < n && dict[num] == 1) {
			return false
		}
		dict[num] += 1
	}
	return true
}
