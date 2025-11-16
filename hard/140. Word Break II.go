package hard

func wordBreak(s string, wordDict []string) []string {
	dict := map[string]bool{}
	for _, str := range wordDict {
		dict[str] = true
	}
	n := len(s)
	stringDict := make([][]string, n)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if dict[s[i:j+1]] {
				stringDict[i] = append(stringDict[i], s[i:j+1])
			}
		}
	}

	ans := []string{}
	curr := ""
	var dfs func(index int)
	dfs = func(index int) {
		if index == n {
			curr = curr[:len(curr)-1]
			ans = append(ans, curr)
			return
		}
		for i := 0; i < len(stringDict[index]); i++ {
			length := len(curr)
			curr += stringDict[index][i] + " "
			dfs(index + len(stringDict[index][i]))
			curr = curr[:length]
		}
	}

	dfs(0)
	return ans
}
