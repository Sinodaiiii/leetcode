package medium

func numTrees(n int) int {
	num := make([]int, n+1)
	num[0], num[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			num[i] += num[j] * num[i-j-1]
		}
	}

	return num[n]
}
