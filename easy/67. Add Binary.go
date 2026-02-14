package easy

func addBinary260215(a string, b string) string {
	n := max(len(a), len(b))
	ia, ib := len(a)-1, len(b)-1
	flag := 0
	ans := ""
	for i := n - 1; i >= 0; i-- {
		ca, cb := 0, 0
		if ia >= 0 {
			ca = int(a[ia] - '0')
			ia -= 1
		}
		if ib >= 0 {
			cb = int(b[ib] - '0')
			ib -= 1
		}
		curr := ca + cb + flag
		flag, curr = curr/2, curr%2
		ans = string(curr+'0') + ans
	}
	if flag != 0 {
		ans = "1" + ans
	}
	return ans
}
