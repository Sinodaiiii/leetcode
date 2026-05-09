package medium

func rotateGrid260509(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	circle := min(m/2, n/2)

	getCoor := func(num, mn, nn int) (int, int) {
		if num < nn-1 {
			return (m - mn) / 2, (n-nn)/2 + num
		} else if num < mn+nn-2 {
			return (m-mn)/2 + num - (nn - 1), n - 1 - (n-nn)/2
		} else if num < mn+2*nn-3 {
			return m - 1 - (m-mn)/2, n - 1 - (n-nn)/2 - (num - (mn - 1) - (nn - 1))
		} else {
			return m - 1 - (m-mn)/2 - (num - (mn - 1) - 2*(nn-1)), (n - nn) / 2
		}
	}

	for i := 0; i < circle; i++ {
		mn, nn := m-2*i, n-2*i
		length := 2*(mn-1) + 2*(nn-1)
		offset := k % length
		for j := 0; j < length; j++ {
			currIndex := j
			x, y := getCoor(j, mn, nn)
			nextX, nextY := getCoor((currIndex+offset)%length, mn, nn)
			tmp := abs(grid[nextX][nextY])
			for grid[x][y] > 0 {
				currIndex = (currIndex - offset + length) % length
				tmp, grid[x][y] = grid[x][y], -tmp
				x, y = getCoor(currIndex, mn, nn)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = -grid[i][j]
		}
	}
	return grid
}

func abs1914(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
