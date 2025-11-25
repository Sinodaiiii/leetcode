package hard

func shortestPalindrome(s string) string {
	n := len(s)
	getPL := func(l, r int) int {
		for l >= 0 && r < n && s[l] == s[r] {
			l -= 1
			r += 1
		}
		return l + 1
	}
	left := 0
	ans := s
	for i := n / 2; i >= 0; i-- {
		left = getPL(i, i+1)
		if left == 0 {
			for j := 2*i - left + 2; j < n; j++ {
				ans = string(s[j]) + ans
			}
			return ans
		}
		left = getPL(i, i)
		if left == 0 {
			for j := 2*i - left + 1; j < n; j++ {
				ans = string(s[j]) + ans
			}
			return ans
		}
	}
	return ans
}
