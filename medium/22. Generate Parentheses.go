package medium

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	cur := make([]byte, 0)
	var dfs func(l, r int)
	dfs = func(l, r int) {
		if l == 0 && r == 0 {
			str := ""
			for i := 0; i < 2*n; i++ {
				str += string(cur[i])
			}
			ans = append(ans, str)
		}
		if r > 0 {
			cur = append(cur, ')')
			dfs(l, r-1)
			cur = cur[:len(cur)-1]
		}
		if l > 0 {
			cur = append(cur, '(')
			dfs(l-1, r+1)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(n, 0)
	return ans
}
