package hot100

func wordBreak(s string, wordDict []string) bool {
	dict := map[string]bool{}
	for i := 0; i < len(wordDict); i++ {
		dict[wordDict[i]] = true
	}
	isOk := make([]bool, len(s)+1)
	isOk[0] = true
	for i := 0; i <= len(s); i++ {
		if isOk[i] && i != len(s) {
			for j := i + 1; j <= len(s); j++ {
				if _, exist := dict[s[i:j]]; exist {
					isOk[j] = true
				}
			}
		}
	}
	return isOk[len(s)]
}
