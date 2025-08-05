package medium

func totalFruit(fruits []int) int {
	ans := -1
	n := len(fruits)
	curr, prev := -1, -1
	pivot, left := 0, 0
	k := 0
	for i := 0; i < n; i++ {
		if fruits[i] != curr && fruits[i] != prev {
			if k < 2 {
				k += 1
				prev, curr = curr, fruits[i]
				pivot = i
			} else {
				prev, curr = curr, fruits[i]
				ans = max(ans, i-left)
				pivot, left = i, pivot
			}
		} else if fruits[i] == prev {
			pivot = i
			prev, curr = curr, prev
			ans = max(ans, i-left+1)
		} else {
			ans = max(ans, i-left+1)
		}
	}
	if ans < 0 {
		return n
	}
	return ans
}
