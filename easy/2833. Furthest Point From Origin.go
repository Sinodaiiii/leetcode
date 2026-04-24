package easy

func furthestDistanceFromOrigin260424(moves string) int {
	l, r, any := 0, 0, 0
	for _, c := range moves {
		switch c {
		case 'L':
			l += 1
		case 'R':
			r += 1
		case '_':
			any += 1
		}
	}
	if l > r {
		return l + any - r
	}
	return r + any - l
}
