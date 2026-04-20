package easy

func maxDistance260420(colors []int) int {
	ans := 0
	n := len(colors)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if colors[i] != colors[j] {
				ans = max(ans, j-i)
			}
		}
	}
	return ans
}
