package medium

func rotatedDigits260502(n int) int {
	ans := 0
	var dfs func(pre, l int, isGood bool)
	badN := []int{0, 1, 8}
	goodN := []int{2, 5, 6, 9}
	dfs = func(pre, l int, isGood bool) {
		if l == 4 {
			if isGood && pre <= n {
				ans += 1
			}
			return
		}
		pre *= 10
		l += 1
		for _, num := range badN {
			dfs(pre+num, l, isGood)
		}
		for _, num := range goodN {
			dfs(pre+num, l, isGood || true)
		}
	}

	dfs(0, 0, false)
	return ans
}
