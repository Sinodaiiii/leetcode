package easy

func makeFancyString(s string) string {
	n := len(s)
	if n <= 2 {
		return s
	}
	curr := byte('0')
	length := 0
	ans := []byte{}
	for i := 0; i < n; i++ {
		if s[i] != curr {
			curr = s[i]
			length = 1
			ans = append(ans, curr)
		} else {
			length += 1
			if length < 3 {
				ans = append(ans, curr)
			}
		}
	}
	return string(ans)
}
