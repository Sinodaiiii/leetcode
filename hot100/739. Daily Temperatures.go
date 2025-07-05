package hot100

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans[i] = 0
		} else {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return ans
}
