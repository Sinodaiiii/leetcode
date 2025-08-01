package medium

func removeKdigits(num string, k int) string {
	n := len(num)
	if n == k {
		return "0"
	}
	stack := make([]byte, 0)
	for i := 0; i < n; i++ {
		for k != 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k -= 1
			if k == 0 {
				break
			}
		}
		stack = append(stack, num[i])
	}
	first := -1
	for i := 0; i < len(stack); i++ {
		if stack[i] != '0' {
			first = i
			break
		}
	}
	if first == -1 {
		return "0"
	} else {
		stack = stack[first:]
		if len(stack) == 0 || len(stack) <= k {
			return "0"
		}
		return string(stack[:len(stack)-k])
	}
}
