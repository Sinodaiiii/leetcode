package medium

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	if n == 1 {
		return 1
	}
	father := make([]int, n)
	for i := 1; i < n; i++ {
		father[i] = i
	}
	ans := n
	var find func(int) int
	var union func(int, int)

	find = func(city int) int {
		if father[city] != city {
			father[city] = find(father[city])
		}
		return father[city]
	}

	union = func(a, b int) {
		ans -= 1
		father[b] = father[a]
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				if find(i) != find(j) {
					union(find(i), find(j))
				}
			}
		}
	}
	return ans
}
