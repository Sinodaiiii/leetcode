package easy

func minimumOperations3190(nums []int) int {
	ans := 0
	for _, num := range nums {
		if num%3 != 0 {
			ans += 1
		}
	}
	return ans
}
