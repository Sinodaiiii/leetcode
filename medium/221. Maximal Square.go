package medium

func convert(byteSlices [][]byte) [][]int {
	intSlices := make([][]int, len(byteSlices))
	for i, byteSlice := range byteSlices {
		intSlice := make([]int, len(byteSlice))
		for j, b := range byteSlice {
			intSlice[j] = int(b) - 48
		}
		intSlices[i] = intSlice
	}
	return intSlices
}

func maximalSquare(matrix [][]byte) int {
	matrixInt := convert(matrix)
	m, n := len(matrixInt), len(matrixInt[0])
	var ans int = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrixInt[i][j] > ans {
				ans = matrixInt[i][j]
				break
			}
		}
	}
	if ans == 0 {
		return 0
	}

	maxLen := max(m, n)
	newAns := true
	for newAns {
		if ans == maxLen {
			break
		}
		newAns = false
		for i := m - ans - 1; i >= 0; i-- {
			for j := n - ans - 1; j >= 0; j-- {
				if matrixInt[i][j] >= ans && matrixInt[i+1][j] >= ans && matrixInt[i][j+1] >= ans && matrixInt[i+1][j+1] >= ans {
					matrixInt[i][j] = ans + 1
					newAns = true
				}
			}
		}
		if newAns {
			ans += 1
		}
	}
	return ans * ans
}
