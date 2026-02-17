package easy

func findDisappearedNumbers260217(nums []int) []int {
	for _, num := range nums {
		index := max(num, -num)
		if nums[index-1] > 0 {
			nums[index-1] = -nums[index-1]
		}
	}
	ans := []int{}
	for i, num := range nums {
		if num > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
