package medium

func findAnagrams(s string, p string) []int {
	dict := map[rune]int{}
	for _, char := range p {
		dict[char] += 1
	}
	runes := []rune(s)
	ans := []int{}
	j, k := len(p), 0
	for i := 0; i < len(s); {
		if v, exist := dict[runes[i]]; exist && v >= 1 {
			k++
			dict[runes[i]]--
			if k == j {
				ans = append(ans, i-j+1)
				k--
				dict[runes[i-k]]++
			}
			i++
		} else {
			if k == 0 {
				i++
			} else {
				dict[runes[i-k]]++
				k--
			}
		}
	}
	return ans
}

//func findAnagrams(s, p string) (ans []int) {
//    sLen, pLen := len(s), len(p)
//    if sLen < pLen {
//        return
//    }
//
//    var sCount, pCount [26]int
//    for i, ch := range p {
//        sCount[s[i]-'a']++
//        pCount[ch-'a']++
//    }
//    if sCount == pCount {
//        ans = append(ans, 0)
//    }
//
//    for i, ch := range s[:sLen-pLen] {
//        sCount[ch-'a']--
//        sCount[s[i+pLen]-'a']++
//        if sCount == pCount {
//            ans = append(ans, i+1)
//        }
//    }
//    return
//}

func findAnagrams260217(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	lp := [26]int{}
	for _, c := range p {
		lp[int(c-'a')] += 1
	}
	sp := [26]int{}
	for i := 0; i < len(p); i++ {
		sp[int(s[i]-'a')] += 1
	}
	ans := []int{}
	if lp == sp {
		ans = append(ans, 0)
	}
	for i := 1; i <= len(s)-len(p); i++ {
		sp[int(s[i-1]-'a')] -= 1
		sp[int(s[i+len(p)-1]-'a')] += 1
		if sp == lp {
			ans = append(ans, i)
		}
	}
	return ans
}
