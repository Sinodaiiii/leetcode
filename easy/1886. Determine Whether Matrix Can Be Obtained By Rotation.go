package easy

func findRotation260322(mat [][]int, target [][]int) bool {
	n := len(mat)
	same := true
	for i := 0; i < n && same; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[i][j] {
				same = false
				break
			}
		}
	}
	if same {
		return same
	} else {
		same = true
	}
	for i := 0; i < n && same; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[j][n-1-i] {
				same = false
				break
			}
		}
	}
	if same {
		return same
	} else {
		same = true
	}
	for i := 0; i < n && same; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[n-1-i][n-1-j] {
				same = false
				break
			}
		}
	}
	if same {
		return same
	} else {
		same = true
	}
	for i := 0; i < n && same; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[n-1-j][i] {
				same = false
				break
			}
		}
	}
	if same {
		return same
	} else {
		same = true
	}
	return false
}
