package medium

func combinationSum3(k int, n int) [][]int {
	matrix := make([][][][][]int, k+1)
	for i := 0; i <= k; i++ {
		matrix[i] = make([][][][]int, n+1)
		for j := 0; j <= n; j++ {
			matrix[i][j] = make([][][]int, 10)
		}
	}
	var sum func(num, last, maxNum int) [][]int
	sum = func(num, last, maxNum int) [][]int {
		if last <= 0 {
			return nil
		}
		if num == 1 {
			if last <= maxNum {
				if matrix[1][last][maxNum] == nil {
					matrix[1][last][maxNum] = [][]int{{last}}
				}
				return matrix[1][last][maxNum]
			}
			return nil
		}
		if matrix[num][last][maxNum] != nil {
			return matrix[num][last][maxNum]
		}
		curr := [][]int{}
		for i := maxNum; i >= num; i-- {
			tmps := sum(num-1, last-i, i-1)
			for _, tmp := range tmps {
				curr = append(curr, append(tmp, i))
			}
		}
		matrix[num][last][maxNum] = curr
		return curr
	}
	sum(k, n, 9)
	return matrix[k][n][9]
}
