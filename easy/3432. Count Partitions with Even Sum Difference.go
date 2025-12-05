package easy

func countPartitions(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return 0
	}
	return len(nums) - 1
}
