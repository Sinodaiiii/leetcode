package hot100

func letterCombinations(digits string) []string {
	list := [][]string{
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
		[]string{"m", "n", "o"},
		[]string{"p", "q", "r", "s"},
		[]string{"t", "u", "v"},
		[]string{"w", "x", "y", "z"},
	}
	var dfs func(index int)
	n := len(digits)
	if n == 0 {
		return []string{}
	}
	curr := ""
	ans := []string{}
	dfs = func(index int) {
		if index == n {
			ans = append(ans, curr)
			return
		}
		for i := 0; i < len(list[digits[index]-'2']); i++ {
			curr += list[digits[index]-'2'][i]
			dfs(index + 1)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0)
	return ans
}
