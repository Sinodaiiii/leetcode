package easy

func kLengthApart(nums []int, k int) bool {
	pre := -1
	for i, num := range nums {
		if num == 1 {
			if pre >= 0 && i-pre-1 < k {
				return false
			}
			pre = i
		}
	}
	return true
}
