package hot100

func findAnagrams(s string, p string) []int {
	ns, np := len(s), len(p)
	if ns < np {
		return []int{}
	}
	dict := map[byte]int{}
	for i := 0; i < np; i++ {
		dict[p[i]] += 1
	}
	curr := len(dict)
	for i := 0; i < np; i++ {
		if dict[s[i]] == 0 {
			curr += 1
		} else if dict[s[i]] == 1 {
			curr -= 1
		}
		dict[s[i]] -= 1
	}
	ans := []int{}
	if curr == 0 {
		ans = append(ans, 0)
	}
	for i := np; i < ns; i++ {
		dict[s[i]] -= 1
		if dict[s[i]] == 0 {
			curr -= 1
		} else if dict[s[i]] == -1 {
			curr += 1
		}
		dict[s[i-np]] += 1
		if dict[s[i-np]] == 1 {
			curr += 1
		} else if dict[s[i-np]] == 0 {
			curr -= 1
		}
		if curr == 0 {
			ans = append(ans, i-np+1)
		}
	}
	return ans
}
