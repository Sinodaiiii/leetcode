package easy

func numberOfSpecialChars260526(word string) int {
	upper := make([]bool, 26)
	lower := make([]bool, 26)
	ans := 0
	for _, s := range word {
		if s >= 'a' && s <= 'z' {
			if upper[s-'a'] && !lower[s-'a'] {
				ans += 1
			}
			lower[s-'a'] = true
		} else if s >= 'A' && s <= 'Z' {
			if lower[s-'A'] && !upper[s-'A'] {
				ans += 1
			}
			upper[s-'A'] = true
		}
	}
	return ans
}
