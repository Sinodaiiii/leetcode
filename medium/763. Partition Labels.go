package medium

func partitionLabels(s string) []int {
	n := len(s)
	dict := map[byte]int{}
	for i := 0; i < n; i++ {
		dict[s[i]] = i
	}
	ans := []int{}
	right := -1
	pre := -1
	for i := 0; i < n; i++ {
		if dict[s[i]] > right {
			right = dict[s[i]]
		}
		if i == right {
			ans = append(ans, i-pre)
			pre = i
		}
	}
	return ans
}
