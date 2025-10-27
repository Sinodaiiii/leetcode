package medium

func numberOfBeams(bank []string) int {
	pre, curr := 0, 0
	ans := 0
	for _, row := range bank {
		curr = 0
		for _, byte := range row {
			if byte == '1' {
				curr += 1
			}
		}
		ans += curr * pre
		if curr != 0 {
			pre = curr
		}
	}
	return ans
}
