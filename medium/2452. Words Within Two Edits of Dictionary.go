package medium

func twoEditWords260422(queries []string, dictionary []string) []string {
	ans := []string{}
	dist := func(x, y string) int {
		diff := 0
		for i := range x {
			if x[i] != y[i] {
				diff += 1
			}
		}
		return diff
	}

	for _, query := range queries {
		for _, str := range dictionary {
			if dist(query, str) <= 2 {
				ans = append(ans, query)
				break
			}
		}
	}
	return ans
}
