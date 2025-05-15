package problems

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
