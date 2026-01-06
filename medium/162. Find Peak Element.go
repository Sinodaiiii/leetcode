package medium

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var find func(l, r int) int
	find = func(l, r int) int {
		if l > r {
			return -1
		}
		mid := (l + r) / 2
		if (mid == 0 && nums[mid+1] < nums[mid]) || (mid == len(nums)-1 && nums[mid-1] < nums[mid]) || (mid > 0 && mid < len(nums)-1 && nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1]) {
			return mid
		}
		left := find(l, mid-1)
		if left != -1 {
			return left
		}
		right := find(mid+1, r)
		if right != -1 {
			return right
		}
		return -1
	}

	ans := find(0, len(nums)-1)
	return ans
}

func findPeakElement260107(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if mid > 0 && nums[mid] > nums[mid-1] && mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			return mid
		} else if mid == 0 && nums[mid] > nums[mid+1] || mid == len(nums)-1 && nums[mid] > nums[mid]-1 {
			return mid
		} else if mid < len(nums)-1 && nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return 0
}
