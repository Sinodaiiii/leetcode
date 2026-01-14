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

func isInterleave260114(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	if s1 == "" || s2 == "" {
		return s3 == s1+s2
	}
	n1, n2, n3 := len(s1), len(s2), len(s3)
	dp := make([][]bool, n2+1)
	for i := 0; i <= n2; i++ {
		dp[i] = make([]bool, n1+1)
	}
	dp[0][0] = true
	for i := 1; i <= n3; i++ {
		for j := n2; j > 0; j-- {
			for k := n1; k > 0; k-- {
				tmp := false
				if s3[i-1] == s1[k-1] {
					tmp = tmp || dp[j][k-1]
				}
				if s3[i-1] == s2[j-1] {
					tmp = tmp || dp[j-1][k]
				}
				dp[j][k] = tmp
			}
			if s3[i-1] == s2[j-1] {
				dp[j][0] = dp[j-1][0]
			}
		}
		for k := n1; k > 0; k-- {
			if s3[i-1] == s1[k-1] {
				dp[0][k] = dp[0][k-1]
			}
		}
		dp[0][0] = false
		// fmt.Println(dp)
	}
	return dp[n2][n1]
}
