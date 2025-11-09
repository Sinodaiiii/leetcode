package easy

func countOperations(num1 int, num2 int) int {
	ans := 0
	m, n := max(num1, num2), min(num1, num2)
	for m != 0 && n != 0 {
		ans += m / n
		if m%n == 0 {
			break
		} else {
			m, n = max(m%n, n), min(m%n, n)
		}
	}
	return ans
}
