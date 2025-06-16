package hot100

func longestPalindrome(s string) string {
	var extendString func(l, r int) (int, int)
	n := len(s)
	extendString = func(l, r int) (int, int) {
		for l >= 0 && r < n && s[l] == s[r] {
			l -= 1
			r += 1
		}
		return l + 1, r - 1
	}
	ansL, ansR := 0, 0
	for i := 0; i < n; i++ {
		for j := i; j <= i+1; j++ {
			l, r := extendString(i, j)
			if r-l > ansR-ansL {
				ansL, ansR = l, r
			}
		}
	}
	return s[ansL : ansR+1]
}
