package medium

// o(n3^2), dp could optimize to o(n1*n2), which n1+n2=n3

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}
	var check func(p1, p2, p3 int) bool
	check = func(p1, p2, p3 int) bool {
		if p3 == n3 {
			return true
		}
		if p1 < n1 && s1[p1] == s3[p3] {
			if check(p1+1, p2, p3+1) == true {
				return true
			}
		}
		if p2 < n2 && s2[p2] == s3[p3] {
			if check(p1, p2+1, p3+1) == true {
				return true
			}
		}
		return false
	}
	return check(0, 0, 0)
}
