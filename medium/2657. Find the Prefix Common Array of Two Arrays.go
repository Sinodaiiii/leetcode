package medium

func findThePrefixCommonArray260520(A []int, B []int) []int {
	n := len(A)
	ans := make([]int, n)
	remain := 0
	dict := make([]int, n+1)
	for i := 0; i < n; i++ {
		dict[A[i]] += 1
		dict[B[i]] -= 1
		remain += 2
		if dict[A[i]] == 0 {
			remain -= 2
		}
		if B[i] != A[i] && dict[B[i]] == 0 {
			remain -= 2
		}
		ans[i] = ((i+1)*2 - remain) / 2
	}
	return ans
}
