package medium

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}
	dict := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		strb := []byte(strs[i])
		sort.Slice(strb, func(i, j int) bool {
			return strb[i] <= strb[j]
		})
		str := string(strb)
		dict[str] = append(dict[str], strs[i])
	}
	ret := make([][]string, 0)
	for _, value := range dict {
		ret = append(ret, value)
	}
	return ret
}
