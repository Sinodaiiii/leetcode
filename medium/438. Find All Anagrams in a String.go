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
