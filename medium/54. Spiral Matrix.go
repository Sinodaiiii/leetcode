package medium

import "math"

func spiralOrder(matrix [][]int) []int {
	status := map[int][2]int{0: {0, 1}, 1: {1, 0}, 2: {0, -1}, 3: {-1, 0}}
	curs, i, j := 0, 0, -1
	m, n := len(matrix), len(matrix[0])
	step := m * n
	ans := make([]int, 0)
	for step > 0 {
		i += status[curs][0]
		j += status[curs][1]
		ans = append(ans, matrix[i][j])
		matrix[i][j] = math.MinInt32
		if i+status[curs][0] >= m || i+status[curs][0] < 0 || j+status[curs][1] >= n || j+status[curs][1] < 0 || matrix[i+status[curs][0]][j+status[curs][1]] == math.MinInt32 {
			curs = (curs + 1) % 4
		}
		step--
	}
	return ans
}
