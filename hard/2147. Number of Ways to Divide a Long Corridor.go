package hard

func numberOfWays(corridor string) int {
	flag := 0
	pNum := 0
	ans := 1
	for _, c := range corridor {
		if c == 'S' {
			if flag == 2 {
				flag = 1
				ans *= pNum + 1
				ans = ans % (1e9 + 7)
			} else {
				flag += 1
			}
			pNum = 0
		} else {
			pNum += 1
		}
	}
	if flag <= 1 {
		return 0
	}
	return ans
}
