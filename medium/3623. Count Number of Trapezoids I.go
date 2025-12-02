package medium

func countTrapezoids(points [][]int) int {
	dict := map[int]int{}
	total := map[int]int{}
	sum := 0
	for _, point := range points {
		dict[point[1]] += 1
		total[point[1]] += dict[point[1]] - 1
		sum += dict[point[1]] - 1
	}
	// fmt.Println(dict)
	// fmt.Println(total)
	// fmt.Println(sum)
	ans := 0
	for _, v := range total {
		// fmt.Println(sum, v)
		ans += (sum - v) * v % (1e9 + 7)
		sum -= v
	}
	return ans % (1e9 + 7)
}
