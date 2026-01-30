package easy

func nextGreatestLetter260131(letters []byte, target byte) byte {
	index := 0
	l, r := 0, len(letters)-1
	for l <= r {
		mid := (l + r) / 2
		if letters[mid] > target {
			index = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return letters[index]
}
