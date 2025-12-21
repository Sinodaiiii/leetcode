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

func findRepeatedDnaSequences2(s string) []string {
	if len(s) < 10 {
		return nil
	}
	dict := map[byte]int{'A': 1, 'C': 2, 'G': 3, 'T': 4}
	num := map[int]int{}
	curr := 0
	for i := 0; i < 10; i++ {
		curr = curr*10 + dict[s[i]]
	}
	num[curr] = 1
	ans := []string{}
	for i := 10; i < len(s); i++ {
		curr = curr%1000000000*10 + dict[s[i]]
		if num[curr] == 1 {
			ans = append(ans, s[i-9:i+1])
		}
		num[curr] += 1
	}
	return ans
}
