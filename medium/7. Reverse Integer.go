package medium

import "math"

func reverse260128(x int) int {
	flag := 1
	if x < 0 {
		x = -x
		flag = -1
	}
	y := 0
	for x >= 10 {
		y += x % 10
		y *= 10
		x = x / 10
	}
	if x+y > int(math.Pow(2, 31)) || (flag == 1 && x+y == int(math.Pow(2, 31))) {
		return 0
	}
	return flag * (x + y)
}
