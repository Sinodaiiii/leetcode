package medium

import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = math.MaxInt32
		}
		matrix[i][i] = 0
	}
	for _, edge := range edges {
		matrix[edge[0]][edge[1]] = edge[2]
		matrix[edge[1]][edge[0]] = edge[2]
	}
	for i := 0; i < n; i++ {
		for s := 0; s < n; s++ {
			for d := 1; d < n; d++ {
				matrix[s][d] = min(matrix[s][d], matrix[s][i]+matrix[i][d])
				matrix[d][s] = matrix[s][d]
			}
		}
	}
	ans, minnum := 0, math.MaxInt32
	for i := n - 1; i >= 0; i-- {
		num := 0
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] <= distanceThreshold {
				num += 1
			}
		}
		if num < minnum {
			ans = i
			minnum = num
		}
	}
	return ans
}
