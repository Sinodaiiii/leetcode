package hard

func calculate(s string) int {
	stack := make([]int, 0)
	flag := 1
	curr := 1
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		} else if s[i] == '+' {
			curr = 1
		} else if s[i] == '-' {
			curr = -1
		} else if s[i] == '(' {
			flag *= curr
			stack = append(stack, curr)
			curr = 1
		} else if s[i] == ')' {
			flag *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			j := i
			num := 0
			for ; j < len(s); j++ {
				if s[j] >= '0' && s[j] <= '9' {
					num = num*10 + int(s[j]-'0')
				} else {
					break
				}
			}
			ans += flag * curr * num
			if j == len(s) {
				break
			}
			i = j - 1
		}
	}
	return ans
}
