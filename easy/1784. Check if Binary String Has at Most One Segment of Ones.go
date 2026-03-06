package easy

func checkOnesSegment260306(s string) bool {
	if s[0] == '0' {
		return false
	}
	pre := false
	for _, c := range s {
		if c == '0' {
			pre = true
		} else {
			if pre {
				return false
			}
		}
	}
	return true
}
