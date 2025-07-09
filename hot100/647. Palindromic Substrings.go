package hot100

func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			ans++
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			ans++
			l--
			r++
		}
	}
	return ans
}
