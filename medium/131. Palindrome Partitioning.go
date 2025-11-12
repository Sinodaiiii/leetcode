package medium

func partition131(s string) [][]string {
	n := len(s)
	pLength := make([][]int, n)
	var palindromeDist func(i, j int) int
	palindromeDist = func(i, j int) int {
		for i >= 0 && j < n && s[i] == s[j] {
			i -= 1
			j += 1
		}
		return i + 1
	}
	for i := 0; i < n; i++ {
		odd := palindromeDist(i, i)
		for j := odd; j <= i; j++ {
			pLength[j] = append(pLength[j], 2*i-j)
		}
		even := palindromeDist(i, i+1)
		for j := even; j <= i; j++ {
			pLength[j] = append(pLength[j], 2*i-j+1)
		}
	}
	ans := [][]string{}
	curr := []string{}
	var dfs func(index int)
	dfs = func(index int) {
		if index == n {
			tmp := make([]string, len(curr))
			copy(tmp, curr)
			ans = append(ans, tmp)
			return
		}
		for i := 0; i < len(pLength[index]); i++ {
			curr = append(curr, s[index:pLength[index][i]+1])
			dfs(pLength[index][i] + 1)
			curr = curr[:len(curr)-1]
		}
	}
	dfs(0)
	return ans
}
