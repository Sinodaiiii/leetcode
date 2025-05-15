package medium

//func uniquePaths(m int, n int) int {
//	m -= 1
//	n -= 1
//	t1 := n + m
//	t2 := m
//	ans := 1
//	ansm := 1
//	for i := 0; i < m; i++ {
//		ans = ans * t1
//		ansm = ansm * t2
//		t1 -= 1
//		t2 -= 1
//	}
//	return ans / ansm
//}

func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		matrix[i][0] = 1
	}
	for i := 0; i < n; i++ {
		matrix[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}
	return matrix[m-1][n-1]
}
