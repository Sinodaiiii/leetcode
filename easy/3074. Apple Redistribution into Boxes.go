package easy

import "sort"

func minimumBoxes260208(apple []int, capacity []int) int {
	total := 0
	for _, a := range apple {
		total += a
	}
	sort.Slice(capacity, func(i, j int) bool { return capacity[i] >= capacity[j] })
	ans := 0
	for _, c := range capacity {
		total -= c
		ans += 1
		if total <= 0 {
			return ans
		}
	}
	return len(capacity)
}
