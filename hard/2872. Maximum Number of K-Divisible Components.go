package hard

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	visited := make([]bool, n)
	nexts := make([][]int, n)
	for _, edge := range edges {
		nexts[edge[0]] = append(nexts[edge[0]], edge[1])
		nexts[edge[1]] = append(nexts[edge[1]], edge[0])
	}
	ans := 1
	var dfs func(index int) int
	dfs = func(index int) int {
		visited[index] = true
		sum := values[index]
		nei := 0
		for _, next := range nexts[index] {
			if visited[next] {
				continue
			}
			nei = dfs(next)
			if nei%k == 0 {
				ans += 1
			} else {
				sum += nei
			}
		}
		return sum % k
	}
	dfs(0)
	return ans
}
