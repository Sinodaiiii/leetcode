package easy

func isHappy(n int) bool {
	dict := map[int]bool{}
	for n != 1 {
		if dict[n] {
			return false
		}
		dict[n] = true
		curr := 0
		for n != 0 {
			tmp := n % 10
			curr += tmp * tmp
			n = n / 10
		}
		n = curr
	}
	return true
}
