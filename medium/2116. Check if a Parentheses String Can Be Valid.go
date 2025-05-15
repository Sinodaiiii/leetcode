package medium

func canBeValid(s string, locked string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if locked[i] == '0' {
			stack = append(stack, ' ')
		} else if s[i] == '(' {
			stack = append(stack, '(')
		} else {
			pare := false
			for len(stack) > 0 {
				if len(stack) >= 2 && stack[len(stack)-1] == ' ' && stack[len(stack)-2] == ' ' {
					stack = stack[:len(stack)-2]
				} else if stack[len(stack)-1] == '(' || stack[len(stack)-1] == ' ' {
					stack = stack[:len(stack)-1]
					pare = true
					break
				}
			}
			if !pare {
				return false
			}
		}
	}
	stack2 := make([]byte, 0)
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] == ' ' {
			stack2 = append(stack2, ' ')
		} else {
			pare := false
			for len(stack2) > 0 {
				if len(stack2) >= 2 && stack2[len(stack2)-1] == ' ' && stack2[len(stack2)-2] == ' ' {
					stack2 = stack2[:len(stack2)-2]
				} else if stack2[len(stack2)-1] == ' ' {
					stack2 = stack2[:len(stack2)-1]
					pare = true
					break
				}
			}
			if !pare {
				return false
			}
		}
	}
	if len(stack2)%2 == 0 {
		return true
	}
	return false
}
