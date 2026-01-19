package medium

func maxSideLength260119(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	for i := 1; i < m; i++ {
		mat[i][0] += mat[i-1][0]
	}
	for i := 1; i < n; i++ {
		mat[0][i] += mat[0][i-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			mat[i][j] += mat[i-1][j] + mat[i][j-1] - mat[i-1][j-1]
		}
	}
	// fmt.Println(mat)
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxD := min(i+1, j+1)
			if maxD <= ans {
				continue
			}
			for d := 1; d <= maxD; d++ {
				l, u, lu := 0, 0, 0
				if j-d >= 0 {
					l = mat[i][j-d]
				}
				if i-d >= 0 {
					u = mat[i-d][j]
				}
				if i-d >= 0 && j-d >= 0 {
					lu = mat[i-d][j-d]
				}
				if mat[i][j]-l-u+lu <= threshold {
					ans = max(ans, d)
					// fmt.Println(i, j, ans, d)
				} else {
					break
				}
			}
		}
	}
	return ans
}
