package medium

func minEatingSpeed(piles []int, h int) int {
	result := -1
	for _, pile := range piles {
		result = max(result, pile)
	}
	check875 := func(k int) bool {
		curr := 0
		for _, pile := range piles {
			curr += pile / k
			if pile%k != 0 {
				curr += 1
			}
			if curr > h {
				return false
			}
		}
		return true
	}
	l, r := 1, result
	for l <= r {
		mid := (l + r) / 2
		if check875(mid) {
			result = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return result
}
