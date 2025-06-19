package search

func findFirstGTE(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left < len(nums) {
		return nums[left]
	}
	return -1
}
