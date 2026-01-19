package easy

func minBitwiseArray260120(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i, num := range nums {
		if num%2 == 0 {
			ans[i] = -1
		} else {
			index := 0
			for num%2 != 0 {
				index += 1
				num = num >> 1
			}
			ans[i] = num<<index + 1<<(index-1) - 1
		}
	}
	return ans
}
