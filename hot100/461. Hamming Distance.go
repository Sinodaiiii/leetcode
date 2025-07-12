package hot100

func hammingDistance(x int, y int) int {
	tmp := x ^ y
	ans := 0
	for tmp != 0 {
		if tmp%2 == 1 {
			ans += 1
		}
		tmp = tmp >> 1
	}
	return ans
}
