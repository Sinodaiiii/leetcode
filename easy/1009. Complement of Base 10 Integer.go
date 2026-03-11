package easy

func bitwiseComplement260311(n int) int {
	i := 1
	for i < n {
		i = i<<1 + 1
	}
	return i - n
}
