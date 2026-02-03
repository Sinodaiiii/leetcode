package medium

func numDecodings260204(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, curr := 1, 1
	for i := 1; i < len(s); i++ {
		next := 0
		if s[i] == '0' {
			if s[i-1] == '0' || s[i-1] > '2' {
				return 0
			}
			next = pre
		} else if s[i] > '6' {
			if s[i-1] == '1' {
				next = pre + curr
			} else {
				next = curr
			}
		} else {
			if s[i-1] > '0' && s[i-1] < '3' {
				next = pre + curr
			} else {
				next = curr
			}
		}
		// fmt.Println(i, pre, curr, next)
		pre, curr = curr, next
	}
	return curr
}
