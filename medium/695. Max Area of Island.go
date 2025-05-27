package medium

//func maxAreaOfIsland(grid [][]int) int {
//    m, n := len(grid), len(grid[0])
//    father := make([][][2]int, m)
//    area := make([][]int, m)
//    for i:=0; i<m; i++ {
//        father[i] = make([][2]int, n)
//        area[i] = make([]int, n)
//    }
//    var find func(x, y int) [2]int
//    var union func (x1, y1, x2, y2 int)
//    find = func(x, y int) [2]int {
//        if father[x][y] != [2]int{x, y} {
//            father[x][y] = find(father[x][y][0], father[x][y][1])
//        }
//        return father[x][y]
//    }
//    union = func (x1, y1, x2, y2 int) {
//        area[x2][y2] += area[father[x1][y1][0]][father[x1][y1][1]]
//        father[father[x1][y1][0]][father[x1][y1][1]] = [2]int{x2, y2}
//        // fmt.Println(x1, y1, x2, y2, area[x2][y2])
//    }
//
//    ans := 0
//    for i:=0; i<m; i++ {
//        for j:=0; j<n; j++ {
//            if grid[i][j] == 1 {
//                father[i][j] = [2]int{i, j}
//                area[i][j] = 1
//                if i>=1 && grid[i-1][j] == 1 && find(i-1, j) != find(i, j) {
//                    union(i-1, j, i, j)
//                }
//                if j>=1 && grid[i][j-1] == 1 && find(i, j-1) != find(i, j) {
//                    union(i, j-1, i, j)
//                }
//                ans = max(ans, area[i][j])
//            }
//        }
//    }
//    return ans
//}

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	step := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	ans := 0
	cur := 0
	var search func(x, y int)
	search = func(x, y int) {
		for c := 0; c < 4; c++ {
			xi, yi := step[c][0], step[c][1]
			if 0 <= x+xi && x+xi < m && 0 <= y+yi && y+yi < n && grid[x+xi][y+yi] == 1 && !visited[x+xi][y+yi] {
				cur += 1
				visited[x+xi][y+yi] = true
				search(x+xi, y+yi)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				cur = 1
				visited[i][j] = true
				search(i, j)
				ans = max(cur, ans)
				// fmt.Println(i, j, cur)
			}
		}
	}
	return ans
}
