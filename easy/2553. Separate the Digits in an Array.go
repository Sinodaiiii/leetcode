package easy

func separateDigits260511(nums []int) []int {
	ans := make([]int, 0, 5*len(nums))
	for _, num := range nums {
		tmp := make([]int, 0, 5)
		for num != 0 {
			tmp = append(tmp, num%10)
			num = num / 10
		}
		for i := len(tmp) - 1; i >= 0; i-- {
			ans = append(ans, tmp[i])
		}
	}
	return ans
}
