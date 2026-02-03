package easy

func isPalindrome260203(s string) bool {
	l, r := 0, len(s)-1
	retLetter := func(c byte) int {
		if c >= 'a' && c <= 'z' {
			return int(c - 'a')
		} else if c >= 'A' && c <= 'Z' {
			return int(c - 'A')
		} else if c >= '0' && c <= '9' {
			return 100 + int(c-'0')
		}
		return -1
	}
	for l < r {
		for l < len(s) && retLetter(s[l]) < 0 {
			l += 1
		}
		for r >= 0 && retLetter(s[r]) < 0 {
			r -= 1
		}
		if l >= r {
			break
		}
		if retLetter(s[l]) != retLetter(s[r]) {
			return false
		}
		l += 1
		r -= 1
	}
	return true
}
