package hard

func largestRectangleArea(heights []int) int {
	cal := make([]int, len(heights))
	ans := heights[0]
	for i := 0; i < len(heights); i++ {
		if heights[i] == cal[i] {
			continue
		}
		cal[i] = heights[i]
		ans = max(ans, heights[i])
		for j := i + 1; j < len(heights); j++ {
			if heights[j] <= cal[j] {
				ans = max(ans, (j-i)*cal[j-1])
				break
			}
			if heights[j] >= cal[j-1] {
				cal[j] = cal[j-1]
			} else {
				ans = max(ans, (j-i)*cal[j-1])
				cal[j] = heights[j]
			}
		}
		ans = max(ans, (len(heights)-i)*cal[len(heights)-1])
	}
	return ans
}

func largestRectangleArea2(heights []int) int {
	stack := []int{-1}
	n := len(heights)
	l := make([]int, n)
	l[0] = -1
	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		l[i] = stack[len(stack)-1]
		stack = append(stack, i)
	}
	stack = []int{n}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 1 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		ans = max(ans, (stack[len(stack)-1]-l[i]-1)*heights[i])
		stack = append(stack, i)
	}
	return ans
}
