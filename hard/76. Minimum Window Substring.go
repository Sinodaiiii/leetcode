package hard

import "math"

func minWindow76(s string, t string) string {
	dict := map[byte]int{}
	for _, ts := range t {
		dict[byte(ts)] += 1
	}
	remain := len(dict)
	l, r := 0, 0
	n := len(s)
	ans := math.MaxInt32
	ansL := -1
	for ; r < n; r++ {
		curr := s[r]
		if _, exist := dict[curr]; !exist {
			continue
		}
		dict[curr] -= 1
		if dict[curr] == 0 {
			remain -= 1
			for ; l <= r && remain == 0; l++ {
				if r-l+1 < ans {
					ans = r - l + 1
					ansL = l
				}
				currl := s[l]
				if _, exist := dict[currl]; exist {
					if dict[currl] == 0 {
						remain += 1
					}
					dict[currl] += 1
				}
			}
		}
	}
	// fmt.Println(ansL, ans)
	if ans <= n {
		return s[ansL : ansL+ans]
	}
	return ""
}
