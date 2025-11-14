package medium

func rangeAddQueries(n int, queries [][]int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	for _, q := range queries {
		for i := q[0]; i <= q[2]; i++ {
			ans[i][q[1]] += 1
			if q[3]+1 < n {
				ans[i][q[3]+1] -= 1
			}
		}
	}
	for i := 0; i < n; i++ {
		curr := 0
		for j := 0; j < n; j++ {
			curr += ans[i][j]
			ans[i][j] = curr
		}
	}
	return ans
}
