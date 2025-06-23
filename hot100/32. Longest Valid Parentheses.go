package hot100

func longestValidParentheses(s string) int {
	n := len(s)
	stack := make([]int, 0)
	ans, curr := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, -1)
		} else {
			if len(stack) > 0 {
				index := len(stack) - 1
				curr = 0
				if stack[index] > 0 {
					curr = stack[index]
					index -= 1
				}
				if index >= 0 && stack[index] == -1 {
					curr += 1
					index -= 1
					if index >= 0 && stack[index] > 0 {
						stack[index] += curr
						stack = stack[:index+1]
						continue
					}
					stack[index+1] = curr
					stack = stack[:index+2]
				} else {
					stack = append(stack, -2)
				}
			} else {
				stack = append(stack, -2)
			}
		}
	}
	for i := 0; i < len(stack); i++ {
		ans = max(ans, stack[i])
	}
	return ans * 2
}
