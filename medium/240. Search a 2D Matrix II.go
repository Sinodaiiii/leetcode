package medium

func searchMatrix240(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for {
		if target == matrix[i][j] {
			return true
		} else if target < matrix[i][j] {
			j -= 1
			if j < 0 {
				return false
			}
		} else {
			i += 1
			if i >= m {
				return false
			}
		}
	}
}
