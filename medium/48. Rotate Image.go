package medium

func rotate(matrix [][]int) {
	n := len(matrix)
	t := (n + 1) / 2
	k := n
	for i := 0; i < t; i++ {
		for j := 0; j < k-1; j++ {
			matrix[i+0+j][i+0], matrix[i+0][n-1-j-i], matrix[n-1-j-i][n-1-i], matrix[n-1-i][i+0+j] = matrix[n-1-i][i+0+j], matrix[i+0+j][i+0], matrix[i+0][n-1-j-i], matrix[n-1-j-i][n-1-i]
		}
		k = (k + 1) / 2
	}
}
