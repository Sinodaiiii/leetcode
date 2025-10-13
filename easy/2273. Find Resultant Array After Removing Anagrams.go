package easy

import "sort"

func removeAnagrams(words []string) []string {
	ans := []string{}
	pre := ""
	for i := 0; i < len(words); i++ {
		currByte := []byte(words[i])
		sort.Slice(currByte, func(i, j int) bool { return currByte[i] <= currByte[j] })
		curr := string(currByte)
		if curr != pre {
			ans = append(ans, words[i])
			pre = curr
		}
	}
	return ans
}
