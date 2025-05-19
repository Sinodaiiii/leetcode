package medium

func reverseWords(s string) string {
	n := len(s)
	stack := make([]string, 0)
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			continue
		}
		l := i
		for i < n && s[i] != ' ' {
			i++
		}
		stack = append(stack, s[l:i])
	}
	n = len(stack)
	ret := stack[n-1]
	for i := n - 2; i >= 0; i-- {
		ret += " " + stack[i]
	}
	return ret
}
