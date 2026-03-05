package easy

func minOperations260305(s string) int {
	pre0, pre1 := 0, 0
	for _, c := range s {
		if c == '0' {
			pre0, pre1 = pre1, pre0+1
		} else {
			pre0, pre1 = pre1+1, pre0
		}
	}
	return min(pre0, pre1)
}
