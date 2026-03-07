package medium

func minFlips260307(s string) int {
	n := len(s)
	l0, l1 := 0, 0
	r0, r1 := 0, 0
	for i := n - 1; i >= 0; i-- {
		if (n-1-i)%2 == 0 {
			if s[i] == '1' {
				r0 += 1
			} else {
				r1 += 1
			}
		} else {
			if s[i] == '0' {
				r0 += 1
			} else {
				r1 += 1
			}
		}
	}
	ans := min(n, l0+r1, l1+r0)
	for i := 0; i < n; i++ {
		if (n-1-i)%2 == 0 {
			if s[i] == '1' {
				r0 -= 1
			} else {
				r1 -= 1
			}
		} else {
			if s[i] == '0' {
				r0 -= 1
			} else {
				r1 -= 1
			}
		}
		if i%2 == 0 {
			if s[i] == '1' {
				l0 += 1
			} else {
				l1 += 1
			}
		} else {
			if s[i] == '0' {
				l0 += 1
			} else {
				l1 += 1
			}
		}
		ans = min(ans, l0+r1, l1+r0)
	}
	return ans
}
