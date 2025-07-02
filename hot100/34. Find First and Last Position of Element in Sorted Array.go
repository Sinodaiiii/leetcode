package hot100

func searchRange(nums []int, target int) []int {
	getLT := func(l, r int) int {
		result := -1
		for l <= r {
			mid := (l + r) / 2
			if nums[mid] < target {
				result = mid
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return result + 1
	}
	getGT := func(l, r int) int {
		result := r + 1
		for l <= r {
			mid := (l + r) / 2
			if nums[mid] > target {
				result = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return result - 1
	}
	l, r := 0, len(nums)-1
	pivot := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			pivot = mid
			break
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if pivot == -1 {
		return []int{-1, -1}
	}
	l = getLT(0, pivot)
	r = getGT(pivot, len(nums)-1)
	return []int{l, r}
}
