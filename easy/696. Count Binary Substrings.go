package easy

func countBinarySubstrings260219(s string) int {
	pre, curr := 0, 1
	ans := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curr += 1
		} else {
			ans += min(pre, curr)
			pre, curr = curr, 1
		}
	}
	return ans + min(pre, curr)
}
