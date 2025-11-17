package hard

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				if dungeon[i][j] >= 0 {
					dungeon[i][j] = 1
				} else {
					dungeon[i][j] = 1 - dungeon[i][j]
				}
				continue
			}
			r, d := math.MaxInt32, math.MaxInt32
			if i != m-1 {
				d = dungeon[i+1][j]
			}
			if j != n-1 {
				r = dungeon[i][j+1]
			}
			next := min(r, d)
			if dungeon[i][j] >= 0 {
				dungeon[i][j] = max(next-dungeon[i][j], 1)
			} else {
				dungeon[i][j] = next - dungeon[i][j]
			}
			// fmt.Println("i: ", i)
			// fmt.Println("j: ", j)
			// fmt.Println(dungeon)
		}
	}
	return dungeon[0][0]
}
