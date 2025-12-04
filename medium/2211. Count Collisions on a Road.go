package medium

func countCollisions(directions string) int {
	ans := 0
	stop := false
	rNum := 0
	for _, d := range directions {
		if d == 'S' {
			ans += rNum
			rNum = 0
			stop = true
		} else if d == 'R' {
			rNum += 1
		} else {
			if rNum != 0 {
				ans += rNum + 1
				rNum = 0
				stop = true
			} else {
				if stop {
					ans += 1
				}
			}
		}
	}

	return ans
}
