package medium

func longestBalanced260212(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		dict := map[byte]int{}
		m := 0
		cNum := 0
		for j := i; j < len(s); j++ {
			dict[s[j]] += 1
			if dict[s[j]] > m {
				m = dict[s[j]]
				cNum = 1
			} else if dict[s[j]] == m {
				cNum += 1
			}
			// fmt.Println(i, j, ans, cNum, m, dict)
			if cNum == len(dict) {
				ans = max(ans, j-i+1)
			}
			// fmt.Println(ans)
		}
	}
	return ans
}
