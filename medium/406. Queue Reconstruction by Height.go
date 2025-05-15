package medium

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})
	l := len(people)
	ans := make([][]int, l)
	for i, person := range people {
		k := person[1]
		j := 0
		for len(ans[j]) != 0 {
			j++
		}
		for k > 0 {
			j++
			for len(ans[j]) != 0 {
				j++
			}
			k--
		}
		ans[j] = people[i]
	}
	return ans
}
