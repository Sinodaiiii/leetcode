package medium

func rotatedDigits260502(n int) int {
	ans := 0
	var dfs func(pre, l int, good bool)
	badN := []int{0, 1, 8}
	goodN := []int{2, 5, 6, 9}
	dfs = func(pre, l int, good bool) {
		if l == 4 {
			if good && pre <= n {
				ans += 1
			}
			return
		}
		pre *= 10
		l += 1
		for _, num := range badN {
			dfs(pre+num, l, good)
		}
		for _, num := range goodN {
			dfs(pre+num, l, good || true)
		}
	}

	dfs(0, 0, false)
	return ans
}
