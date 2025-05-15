package problems

func longestCycle(edges []int) int {
	ans, t := -1, 1
	visited := make([]int, len(edges))
	for i := 0; i < len(edges); i++ {
		if visited[i] != 0 {
			continue
		}
		j := i
		for j != -1 && visited[j] == 0 {
			visited[j] = t
			j = edges[j]
		}
		if j != -1 && visited[j] == t {
			l, k := 1, j
			j = edges[j]
			for j != k {
				l++
				j = edges[j]
			}
			ans = max(ans, l)
		}
		t++
	}
	return ans
}
