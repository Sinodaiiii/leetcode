package medium

import "strconv"

func countAndSay(n int) string {
	ans := "1"
	for i := 2; i <= n; i++ {
		na := len(ans)
		str := ""
		count := 1
		curr := ans[na-1]
		j := na - 1
		for {
			j--
			if j < 0 || ans[j] != curr {
				str = strconv.Itoa(count) + string(curr) + str
				if j < 0 {
					break
				}
				count = 1
				curr = ans[j]
			} else {
				count++
			}
		}
		ans = str
	}
	return ans
}
