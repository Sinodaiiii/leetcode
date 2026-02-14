package medium

func champagneTower260214(poured int, query_row int, query_glass int) float64 {
	if query_row == 0 && query_glass == 0 {
		return min(float64(poured), 1.0)
	}
	type node struct {
		i int
		j int
	}
	dict := map[node]float64{node{0, 0}: float64(poured) - 1.0}
	var cp func(row, glass int) float64
	cp = func(row, glass int) float64 {
		if ret, exist := dict[node{row, glass}]; exist {
			return ret
		}
		if glass < 0 || glass > row {
			return 0.0
		}
		l, r := cp(row-1, glass-1), cp(row-1, glass)
		next := l/2 + r/2 - 1
		if next < 0 {
			next = 0
		}
		dict[node{row, glass}] = next
		return next
	}

	remain := (cp(query_row-1, query_glass-1) + cp(query_row-1, query_glass)) / 2
	if remain >= 1.0 {
		return 1.0
	} else if remain < 0.0 {
		return 0.0
	}
	return remain
}
