package medium

import "math"

type island struct {
	id     int
	coords [][2]int
}

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	islands := make([]island, 0)
	tag := make([][]int, m)
	for i := 0; i < m; i++ {
		tag[i] = make([]int, n)
		for j := 0; j < n; j++ {
			tag[i][j] = math.MinInt32
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				newIsland := true
				if i-1 >= 0 && tag[i-1][j] >= 0 { //上侧有岛
					newIsland = false
					tag[i][j] = tag[i-1][j]
					islands[tag[i][j]].coords = append(islands[tag[i][j]].coords, [2]int{i, j})
				}
				if j-1 >= 0 && tag[i][j-1] >= 0 { //左侧有岛
					newIsland = false
					if tag[i][j] < 0 { //自己是独立的
						tag[i][j] = tag[i][j-1]
						islands[tag[i][j]].coords = append(islands[tag[i][j]].coords, [2]int{i, j})
					} else { //自己不是独立的，合并
						if tag[i][j] == tag[i][j-1] {
							continue
						}
						s, t := max(tag[i][j], tag[i][j-1]), min(tag[i][j], tag[i][j-1])
						for _, coord := range islands[s].coords {
							tag[coord[0]][coord[1]] = islands[t].id
						}
						islands[t].coords = append(islands[t].coords, islands[s].coords...)
						islands[s].coords = make([][2]int, 0)
					}
				}

				if newIsland {
					id := len(islands)
					islands = append(islands, island{id, [][2]int{{i, j}}})
					tag[i][j] = id
				}

			}
		}
	}
	ans := 0
	for _, i := range islands {
		if len(i.coords) != 0 {
			ans++
		}
	}
	return ans
}

//func numIslands(grid [][]byte) int {
//	m, n := len(grid), len(grid[0])
//	islands := make([][]int, m)
//	for i := 0; i < m; i++ {
//		islands[i] = make([]int, n)
//		for j := 0; j < n; j++ {
//			islands[i][j] = math.MinInt32
//		}
//	}
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if i == 0 && j == 0 {
//				islands[0][0] = 0
//			}
//			if grid[i][j] != '0' {
//				if i-1 >= 0 && islands[i-1][j] >= 0 {
//					islands[i][j] = islands[i-1][j]
//				}
//			}
//		}
//	}
//
//	ans := 0
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//
//		}
//	}
//	return ans
//}
