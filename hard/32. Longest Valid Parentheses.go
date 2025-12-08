package hard

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	ans := 0
	stack := []int{}
	for _, value := range s {
		if value == '(' {
			stack = append(stack, -1)
		} else if value == ')' {
			for i := len(stack) - 1; i >= 0; i-- {
				if stack[i] == -1 {
					l := 2
					for j := i; j < len(stack); j++ {
						if stack[j] > 0 {
							l += stack[j]
						}
					}
					stack[i] = l
					if i > 0 && stack[i-1] >= 0 {
						stack[i-1] += l
						i -= 1
					}
					stack = stack[:i+1]
					break
				}
				if i == 0 {
					for _, value := range stack {
						ans = max(ans, value)
					}
					stack = []int{}
				}
			}
		}
	}
	for _, value := range stack {
		ans = max(ans, value)
	}
	return ans
}

func longestValidParentheses2(s string) int {
	stack := []int{-1}
	ans := 0
	for index, b := range s {
		if b == ')' && len(stack) > 1 && s[stack[len(stack)-1]] == '(' {
			ans = max(ans, index-stack[len(stack)-2])
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, index)
		}
	}
	return ans
}
