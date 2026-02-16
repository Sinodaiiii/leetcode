package easy

func reverseBits260216(n int) int {
	ans := 0
	for i := 32; i > 0; i-- {
		ans = ans*2 + n%2
		n = n / 2
	}
	return ans
}
