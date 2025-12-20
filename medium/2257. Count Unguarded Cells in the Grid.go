package medium

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	gird := make([][][4]bool, m)
	for i := 0; i < m; i++ {
		gird[i] = make([][4]bool, n)
	}
	for _, wall := range walls {
		gird[wall[0]][wall[1]] = [4]bool{true, true, true, true}
	}
	direction := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	queue := make([][3]int, 0)
	for _, guard := range guards {
		x, y := guard[0], guard[1]
		gird[x][y] = [4]bool{true, true, true, true}
		for i := 0; i < 4; i++ {
			xi, yi := x+direction[i][0], y+direction[i][1]
			if xi >= 0 && xi < m && yi >= 0 && yi < n && !gird[xi][yi][i] {
				queue = append(queue, [3]int{xi, yi, i})
			}
		}
	}
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		gird[curr[0]][curr[1]][curr[2]] = true
		xi, yi := curr[0]+direction[curr[2]][0], curr[1]+direction[curr[2]][1]
		if xi >= 0 && xi < m && yi >= 0 && yi < n && !gird[xi][yi][curr[2]] {
			queue = append(queue, [3]int{xi, yi, curr[2]})
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !gird[i][j][0] && !gird[i][j][1] && !gird[i][j][2] && !gird[i][j][3] {
				ans += 1
			}
		}
	}
	return ans
}
