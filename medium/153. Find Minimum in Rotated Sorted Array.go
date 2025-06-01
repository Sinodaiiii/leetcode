package medium

func findMin(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	l, r := 0, n-1
	for {
		mid := (l + r) / 2
		if nums[mid] < nums[(mid-1+n)%n] {
			return nums[mid]
		} else if nums[mid] < nums[r] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
}
