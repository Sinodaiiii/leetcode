package hot100

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	var match func(m, n int) bool
	var checkP func(n int) bool
	checkP = func(n int) bool {
		if n%2 == 0 {
			return false
		}
		for n >= 0 && p[n] == '*' {
			n -= 2
		}
		return n == -1
	}
	match = func(m, n int) bool {
		if m == -1 && n == -1 {
			return true
		} else if m == -1 {
			return checkP(n)
		} else if n == -1 {
			return false
		}
		if p[n] != '*' {
			if s[m] == p[n] || p[n] == '.' {
				return match(m-1, n-1)
			} else {
				return false
			}
		} else {
			for m < 0 || s[m] == p[n-1] || p[n-1] == '.' {
				if m < -1 {
					return false
				} else if match(m, n-2) == true {
					return true
				} else {
					m -= 1
				}
			}
			return match(m, n-2)
		}
	}

	return match(m-1, n-1)
}
