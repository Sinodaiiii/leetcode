package medium

import "math"

func maxMatrixSum260105(matrix [][]int) int64 {
	m, n := len(matrix), len(matrix[0])
	negNum := 0
	sum := 0
	minNum := math.MaxInt32
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] <= 0 {
				negNum += 1
				sum += -matrix[i][j]
				minNum = min(minNum, -matrix[i][j])
			} else {
				sum += matrix[i][j]
				minNum = min(minNum, matrix[i][j])
			}
		}
	}
	if negNum%2 == 1 {
		return int64(sum - 2*minNum)
	}
	return int64(sum)
}
