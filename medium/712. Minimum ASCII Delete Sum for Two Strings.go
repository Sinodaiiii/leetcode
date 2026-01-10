package medium

func minimumDeleteSum260110(s1 string, s2 string) int {
	n1, n2 := len(s1), len(s2)
	pre := make([]int, n2+1)
	for i := 1; i <= n2; i++ {
		pre[i] = pre[i-1] + int(s2[i-1])
	}
	for i := 1; i <= n1; i++ {
		curr := make([]int, n2+1)
		curr[0] = pre[0] + int(s1[i-1])
		for j := 1; j <= n2; j++ {
			if s1[i-1] == s2[j-1] {
				curr[j] = pre[j-1]
			} else {
				curr[j] = min(pre[j]+int(s1[i-1]), curr[j-1]+int(s2[j-1]))
			}
		}
		pre = curr
	}
	return pre[n2]
}
