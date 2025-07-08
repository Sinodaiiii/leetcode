package hot100

func maximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	width := make([][]int, m)
	for i := 0; i < m; i++ {
		width[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if j == 0 {
					width[i][j] = 1
				} else {
					width[i][j] = width[i][j-1] + 1
				}
			}
		}
	}
	ans := 0
	maxRect := func(column int) {
		left, right := make([]int, m), make([]int, m)
		stack := []int{-1}
		for i := 0; i < m; i++ {
			for len(stack) > 1 && width[i][column] <= width[stack[len(stack)-1]][column] {
				stack = stack[:len(stack)-1]
			}
			// if len(stack) > 1 && width[i][column] == width[stack[len(stack)-1]][column] {
			//     left[i] = stack[len(stack)-2]
			// } else {
			left[i] = stack[len(stack)-1]
			stack = append(stack, i)
			// }
		}
		stack = []int{m}
		for i := m - 1; i >= 0; i-- {
			for len(stack) > 1 && width[i][column] <= width[stack[len(stack)-1]][column] {
				stack = stack[:len(stack)-1]
			}
			// if len(stack) > 1 && width[i][column] == width[stack[len(stack)-1]][column] {
			//     right[i] = stack[len(stack)-2]
			// } else {
			right[i] = stack[len(stack)-1]
			stack = append(stack, i)
			// }
		}
		for i := 0; i < m; i++ {
			ans = max(ans, (right[i]-left[i]-1)*width[i][column])
		}
	}
	for i := 0; i < n; i++ {
		maxRect(i)
	}
	return ans
}
