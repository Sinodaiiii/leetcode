package medium

func searchMatrix(matrix [][]int, target int) bool {
	if target < matrix[0][0] {
		return false
	} else if target == matrix[0][0] {
		return true
	}
	n, m := len(matrix), len(matrix[0])
	l, r := 0, n-1
	for l < r {
		mid := (l + r) / 2
		if matrix[mid][0] == target {
			return true
		}
		if matrix[l][0] > target {
			l -= 1
			break
		}
		if matrix[r][0] < target {
			l = r
			break
		}
		if matrix[mid][0] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	t := l
	l, r = 0, m-1
	for l < r {
		mid := (l + r) / 2
		if matrix[t][mid] == target {
			return true
		}
		if matrix[t][l] > target || matrix[t][r] < target {
			return false
		}
		if matrix[t][mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if matrix[t][l] == target {
		return true
	}
	return false
}

func searchMatrix260130(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1
	row := 0
	for l <= r {
		mid := (l + r) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			row = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	l, r = 0, len(matrix[0])-1
	for l <= r {
		mid := (l + r) / 2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
