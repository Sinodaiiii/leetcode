package medium

func productExceptSelf(nums []int) []int {
	n := len(nums)
	var product int64 = 1
	zeroi := -1
	for i, value := range nums {
		if value == 0 {
			if zeroi >= 0 {
				return make([]int, n)
			}
			zeroi = i
		} else {
			product *= int64(value)
		}
	}
	ans := make([]int, n)
	if zeroi >= 0 {
		ans[zeroi] = int(product)
	} else {
		for i := range nums {
			ans[i] = int(product / int64(nums[i]))
		}
	}
	return ans
}
