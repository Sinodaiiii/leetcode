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

func lengthOfLongestSubstring2(s string) int {
	dict := make([]int, 128)
	for i := 0; i < 128; i++ {
		dict[i] = -1
	}
	l, r := 0, 0
	ans := 0
	for r < len(s) {
		if dict[int(s[r])] < l {
			ans = max(r-l+1, ans)
		} else {
			l = dict[int(s[r])] + 1
		}
		dict[int(s[r])] = r
		r += 1
	}
	return ans
}
