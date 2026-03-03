package medium

func findKthBit260303(n int, k int) byte {
	l := make([]int, n+1)
	l[1] = 1
	for i := 2; i <= n; i++ {
		l[i] = 2*l[i-1] + 1
	}
	turn := false
	for ; n > 1; n-- {
		mid := l[n]/2 + 1
		if k == mid {
			if turn {
				return '0'
			} else {
				return '1'
			}
		} else if k > mid {
			k = 2*mid - k
			turn = !turn
		}
	}
	if turn {
		return '1'
	}
	return '0'
}
