package medium

func countPalindromicSubsequence(s string) int {
	dict := map[rune]*[2]int{}
	for i, c := range s {
		if _, exist := dict[c]; exist {
			dict[c][1] = i
		} else {
			dict[c] = &[2]int{i, i}
		}
	}
	ans := 0
	for _, v := range dict {
		if v[1] > v[0]+1 {
			diff := map[rune]bool{}
			for i := v[0] + 1; i < v[1]; i++ {
				diff[rune(s[i])] = true
			}
			ans += len(diff)
		}
	}
	return ans
}
