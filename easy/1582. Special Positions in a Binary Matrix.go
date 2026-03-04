package easy

func numSpecial260304(mat [][]int) int {
	r := make([]int, len(mat))
	ans := 0
	for i := 0; i < len(mat); i++ {
		curr := -1
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				if curr != -1 {
					curr = -1
					break
				}
				curr = j
			}
		}
		r[i] = curr
	}
	for j := 0; j < len(mat[0]); j++ {
		curr := -1
		for i := 0; i < len(mat); i++ {
			if mat[i][j] == 1 {
				if curr != -1 {
					curr = -1
					break
				}
				curr = i
			}
		}
		if curr >= 0 && r[curr] == j {
			ans += 1
		}
	}
	return ans
}
