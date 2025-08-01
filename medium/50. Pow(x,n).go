package medium

func myPow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if x == 1 || n == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	half := myPow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}
