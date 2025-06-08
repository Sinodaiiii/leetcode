package hard

func trap(height []int) int {
	n := len(height)
	l := make([]int, n)
	l[0] = height[0]
	for i := 1; i < n; i++ {
		l[i] = max(l[i-1], height[i])
	}
	ans := 0
	r := height[n-1]
	for i := n - 2; i >= 1; i-- {
		cur := min(l[i-1], r) - height[i]
		if cur > 0 {
			ans += cur
		}
		r = max(r, height[i])
	}
	return ans
}
