package medium

func countCoveredBuildings(n int, buildings [][]int) int {
	row := map[int]*[2]int{}
	column := map[int]*[2]int{}
	for _, building := range buildings {
		if row[building[0]] == nil {
			row[building[0]] = &[2]int{building[1], building[1]}
		} else {
			row[building[0]][0] = min(row[building[0]][0], building[1])
			row[building[0]][1] = max(row[building[0]][1], building[1])
		}
		if column[building[1]] == nil {
			column[building[1]] = &[2]int{building[0], building[0]}
		} else {
			column[building[1]][0] = min(column[building[1]][0], building[0])
			column[building[1]][1] = max(column[building[1]][1], building[0])
		}
	}
	ans := 0
	for _, building := range buildings {
		if building[0] > column[building[1]][0] && building[0] < column[building[1]][1] {
			if building[1] > row[building[0]][0] && building[1] < row[building[0]][1] {
				ans += 1
			}
		}
	}
	return ans
}
