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

func findFirstLTE(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] <= target {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
