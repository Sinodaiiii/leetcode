package hard

func findSubstring(s string, words []string) []int {
	n, nw, nws := len(s), len(words), len(words[0])
	addW := func(dict map[string]int, input string) {
		dict[input] -= 1
		if dict[input] == 0 {
			delete(dict, input)
		}
	}
	deleteW := func(dict map[string]int, input string) {
		dict[input] += 1
		if dict[input] == 0 {
			delete(dict, input)
		}
	}
	ans := []int{}
	for i := 0; i < nws; i++ {
		dict := map[string]int{}
		for j := 0; j < nw; j++ {
			dict[words[j]] += 1
		}
		left, right := i, i+nw*nws
		if right > n {
			break
		}
		for j := left; j+nws <= right; j += nws {
			addW(dict, s[j:j+nws])
		}
		if len(dict) == 0 {
			ans = append(ans, left)
		}
		for right+nws <= n {
			deleteW(dict, s[left:left+nws])
			left += nws
			addW(dict, s[right:right+nws])
			right += nws
			if len(dict) == 0 {
				ans = append(ans, left)
			}
		}
	}
	return ans
}
