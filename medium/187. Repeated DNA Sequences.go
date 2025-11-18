package medium

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}
	dict := map[string]int{s[:10]: 1}
	ans := []string{}
	for i := 10; i < len(s); i++ {
		if dict[s[i-9:i+1]] == 1 {
			ans = append(ans, s[i-9:i+1])
		}
		dict[s[i-9:i+1]] += 1
	}
	return ans
}
