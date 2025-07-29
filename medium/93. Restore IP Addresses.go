package medium

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	ans := []string{}
	checkAns := func(i, j, k int) {
		str := []string{s[:i+1], s[i+1 : j+1], s[j+1 : k+1], s[k+1:]}
		for _, ss := range str {
			if len(ss) > 3 || len(ss) == 0 {
				return
			}
			tmp, _ := strconv.Atoi(ss)
			if tmp > 255 || strconv.Itoa(tmp) != ss {
				return
			}
		}
		ans = append(ans, fmt.Sprintf("%s.%s.%s.%s", str[0], str[1], str[2], str[3]))
	}

	n := len(s)
	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			for k := j + 1; k < n-1; k++ {
				checkAns(i, j, k)
			}
		}
	}

	return ans
}
