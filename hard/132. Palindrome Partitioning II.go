package hard

func minCut(s string) int {
	n := len(s)
	pLength := make([][]int, n)
	var palindromeDist func(i, j int) int
	palindromeDist = func(i, j int) int {
		for i >= 0 && j < n && s[i] == s[j] {
			i -= 1
			j += 1
		}
		return i + 1
	}
	for i := 0; i < n; i++ {
		odd := palindromeDist(i, i)
		for j := odd; j <= i; j++ {
			pLength[j] = append(pLength[j], 2*i-j+1)
		}
		even := palindromeDist(i, i+1)
		for j := even; j <= i; j++ {
			pLength[j] = append(pLength[j], 2*i-j+2)
		}
	}

	visited := make([]bool, n)
	visited[0] = true
	queue := []int{0}
	curr := -1
	for {
		curr += 1
		length := len(queue)
		for i := 0; i < length; i++ {
			index := queue[i]
			for _, end := range pLength[index] {
				if end == n {
					return curr
				}
				if !visited[end] {
					visited[end] = true
					queue = append(queue, end)
				}
			}
		}
		queue = queue[length:]
	}
	return curr
}
