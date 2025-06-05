package medium

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, column := [4]int{0, 0, 0, 0}, [4]int{0, 0, 0, 0}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				irow, jrow := i/60, i%60
				icolumn, jcolumn := j/60, j%60
				row[irow] = row[irow] | (1 << (jrow + 1))
				column[icolumn] = column[icolumn] | (1 << (jcolumn + 1))
			}
		}
	}
	for i := 0; i < m; i++ {
		irow, jrow := i/60, i%60
		if row[irow]&(1<<(jrow+1)) != 0 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 0; j < n; j++ {
		icolumn, jcolumn := j/60, j%60
		if column[icolumn]&(1<<(jcolumn+1)) != 0 {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
	return
}
