package easy

func mirrorDistance260418(n int) int {
	tmp := n
	re := 0
	for tmp != 0 {
		re = re*10 + tmp%10
		tmp = tmp / 10
	}
	if n-re >= 0 {
		return n - re
	}
	return re - n
}
