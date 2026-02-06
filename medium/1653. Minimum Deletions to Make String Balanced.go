package medium

func minimumDeletions260207(s string) int {
	n, na, nb := len(s), 0, 0
	for _, c := range s {
		if c == 'a' {
			na += 1
		} else {
			nb += 1
		}
	}
	ans := na
	pa := 0
	for _, c := range s {
		if c == 'a' {
			pa += 1
		} else {
			nb -= 1
		}
		ans = min(ans, na-pa+(n-nb-na))
	}
	return ans
}
