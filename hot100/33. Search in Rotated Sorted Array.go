package hot100

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[l] {
			if nums[l] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[r] >= target && nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	if nums[l] != target {
		return -1
	}
	return l
}
