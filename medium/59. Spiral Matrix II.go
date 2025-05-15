package medium

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	if n == 1 {
		ans[0][0] = 1
		return ans
	}
	status := 1
	num, i, j := 1, 0, 0
	for num <= n*n {
		if ans[i][j] != 0 {
			break
		}
		ans[i][j] = num
		num++
		switch status {
		case 1:
			if j == n-1 || ans[i][j+1] != 0 {
				status = 2
				i++
			} else {
				j++
			}
		case 2:
			if i == n-1 || ans[i+1][j] != 0 {
				status = 3
				j--
			} else {
				i++
			}
		case 3:
			if j == 0 || ans[i][j-1] != 0 {
				status = 4
				i--
			} else {
				j--
			}
		case 4:
			if ans[i-1][j] != 0 {
				status = 1
				j++
			} else {
				i--
			}
		}
	}
	return ans
}
