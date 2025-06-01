package medium

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	n := len(tokens)
	for i := 0; i < n; i++ {
		if len(tokens[i]) > 1 || (tokens[i][0] <= '9' && tokens[i][0] >= '0') {
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
		} else {
			ns := len(stack)
			a, b, result := stack[ns-2], stack[ns-1], 0
			if tokens[i][0] == '+' {
				result = a + b
			} else if tokens[i][0] == '-' {
				result = a - b
			} else if tokens[i][0] == '*' {
				result = a * b
			} else if tokens[i][0] == '/' {
				result = a / b
			}
			stack[ns-2] = result
			stack = stack[0 : ns-1]
		}
	}
	return stack[0]
}
