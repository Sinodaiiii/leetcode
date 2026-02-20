package hard

func removeInvalidParentheses260220(s string) []string {
	check := func(curr string) bool {
		lnum := 0
		for _, c := range curr {
			switch c {
			case '(':
				lnum += 1
			case ')':
				if lnum > 0 {
					lnum -= 1
				} else {
					return false
				}
			default:
				continue
			}
		}
		return lnum == 0
	}

	lnum, rnum := 0, 0
	for _, c := range s {
		switch c {
		case '(':
			lnum += 1
		case ')':
			if lnum > 0 {
				lnum -= 1
			} else {
				rnum += 1
			}
		default:
			continue
		}
	}

	ans := []string{}
	n := len(s) - lnum - rnum
	var dfs func(l, r, index int, curr string)
	dfs = func(l, r, index int, curr string) {
		if len(curr) == n && check(curr) {
			ans = append(ans, curr)
		}
		for i := index; i < len(curr); i++ {
			if (curr[i] >= 'a' && curr[i] <= 'z') || (i > 0 && curr[i] == curr[i-1]) {
				continue
			}
			if curr[i] == '(' && l > 0 {
				dfs(l-1, r, i, curr[:i]+curr[i+1:])
			}
			if curr[i] == ')' && r > 0 {
				dfs(l, r-1, i, curr[:i]+curr[i+1:])
			}
		}
	}

	dfs(lnum, rnum, 0, s)
	return ans
}
