package medium

func grayCode(n int) []int {
	ans := make([]int, 1<<n)
	ans[0], ans[1] = 0, 1
	for i := 2; i <= n; i++ {
		l := 1 << (i - 1)
		for i := 0; i < l; i++ {
			ans[i] *= 2
			ans[l*2-1-i] = ans[i] + 1
		}
	}
	return ans
}
