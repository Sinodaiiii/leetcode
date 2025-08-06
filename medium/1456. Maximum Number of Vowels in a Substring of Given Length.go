package medium

func maxVowels(s string, k int) int {
	ans := 0
	dict := map[byte]bool{'a': true, 'o': true, 'e': true, 'i': true, 'u': true}
	for i := 0; i < k; i++ {
		if dict[s[i]] {
			ans += 1
		}
	}
	left, right := 0, k-1
	n := len(s)
	curr := ans
	for right+1 < n {
		if dict[s[left]] {
			curr -= 1
		}
		left += 1
		right += 1
		if dict[s[right]] {
			curr += 1
		}
		ans = max(ans, curr)
	}
	return ans
}
