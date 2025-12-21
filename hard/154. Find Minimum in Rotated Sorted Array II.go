package hard

func findMin(nums []int) int {
	target := 0
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < nums[target] {
			target = mid
		}
		if nums[l] == nums[mid] {
			l += 1
		} else if nums[mid] == nums[r] {
			r -= 1
		} else if nums[mid] < nums[r] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return nums[target]
}
