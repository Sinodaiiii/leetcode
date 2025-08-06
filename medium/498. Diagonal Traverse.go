package medium

func findDiagonalOrder(mat [][]int) []int {
	direction := 1
	i, j := 0, 0
	m, n := len(mat), len(mat[0])
	ans := []int{}
	for i != m-1 || j != n-1 {
		ans = append(ans, mat[i][j])
		if direction == 1 {
			i += -1
			j += 1
		} else {
			i += 1
			j += -1
		}
		if i < 0 || i >= m || j < 0 || j >= n {
			if direction == 1 {
				if j < n {
					i = 0
				} else {
					i += 2
					j -= 1
				}
			} else {
				if i < m {
					j = 0
				} else {
					j += 2
					i -= 1
				}
			}
			direction *= -1
		}

	}
	ans = append(ans, mat[i][j])
	return ans
}
