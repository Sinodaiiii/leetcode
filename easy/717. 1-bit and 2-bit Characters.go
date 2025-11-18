package easy

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	if bits[n-1] != 0 {
		return false
	}
	for i := 0; i < n; i++ {
		if bits[i] == 0 {
			if i == n-1 {
				return true
			} else {
				continue
			}
		} else {
			i += 1
		}
	}
	return false
}
