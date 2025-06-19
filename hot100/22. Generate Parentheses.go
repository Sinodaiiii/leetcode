package hot100

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	curr := ""
	var dfs func(l, r int)
	dfs = func(l, r int) {
		if l == 0 && r == 0 {
			ans = append(ans, curr)
		}
		if l > 0 {
			curr += "("
			dfs(l-1, r+1)
			curr = curr[:len(curr)-1]
		}
		if r > 0 {
			curr += ")"
			dfs(l, r-1)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(n, 0)
	return ans
}
