package medium

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	ans := 1
	dict := map[byte]int{}
	curr := 0
	for i := 0; i < n; i++ {
		if index, exist := dict[s[i]]; !exist {
			curr += 1
		} else {
			curr = min(i-index, curr+1)
		}
		dict[s[i]] = i
		ans = max(ans, curr)
	}
	return ans
}
