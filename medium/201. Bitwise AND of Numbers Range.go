package medium

func rangeBitwiseAnd(left int, right int) int {
	b := make([]int, 32)
	for i := 0; i < 32; i++ {
		curr := 1 << i
		mod := (right + 1) % (curr * 2)
		b[i] += (right + 1) / (curr * 2) * curr
		if mod > curr {
			b[i] += mod - curr
		}
	}
	// fmt.Println(b)
	for i := 0; i < 32; i++ {
		curr := 1 << i
		if left/curr == 0 {
			break
		}
		mod := left % (curr * 2)
		b[i] -= left / (curr * 2) * curr
		if mod > curr {
			b[i] -= mod - curr
		}
	}
	// fmt.Println(b)
	ans := 0
	target := right - left + 1
	for i := 0; i < 32; i++ {
		if b[i] > 0 && b[i] == target {
			ans += 1 << i
		}
	}
	return ans
}
