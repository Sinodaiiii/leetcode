package easy

func countTriples(n int) int {
	dict := map[int]int{}
	ans := 0
	for i := 1; i <= n; i++ {
		ans += dict[i*i] * 2
		for j := 1; j < i; j++ {
			dict[j*j+i*i] += 1
		}
	}
	return ans
}
