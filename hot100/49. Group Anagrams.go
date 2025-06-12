package hot100

import "sort"

func groupAnagrams(strs []string) [][]string {
	dict := map[string]int{}
	ans := make([][]string, 0)
	n := len(strs)
	indexAns := 0
	for i := 0; i < n; i++ {
		curr := []byte(strs[i])
		sort.Slice(curr, func(i, j int) bool { return curr[i] <= curr[j] })
		if index, exist := dict[string(curr)]; exist {
			ans[index] = append(ans[index], strs[i])
		} else {
			dict[string(curr)] = indexAns
			ans = append(ans, []string{strs[i]})
			indexAns += 1
		}
	}
	return ans
}
