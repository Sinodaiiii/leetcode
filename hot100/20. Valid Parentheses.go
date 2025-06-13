package hot100

func isValid(s string) bool {
	n := len(s)
	stack := []byte{}
	pair := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	for i := 0; i < n; i++ {
		if pair[s[i]] != 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pair[s[i]] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
