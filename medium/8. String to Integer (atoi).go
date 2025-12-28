package medium

func myAtoi(s string) int {
	ans := 0
	flag := 1
	mod := 1<<31 - 1
	stage := 0
	for _, c := range s {
		if stage == 0 && c == ' ' {
			continue
		} else if stage == 0 && c == '+' {
			flag = 1
			stage = 1
		} else if stage == 0 && c == '-' {
			flag = -1
			stage = 1
		} else if c >= '0' && c <= '9' {
			stage = 1
			ans = ans*10 + int(c-'0')
			if ans > mod {
				if flag == 1 {
					return mod
				}
				return -mod - 1
			}
		} else {
			break
		}
	}
	if ans > mod {
		if flag == 1 {
			return mod
		}
		return -mod - 1
	}
	return ans * flag
}
