package medium

func countPermutations(complexity []int) int {
	ans := 1
	mod := int(1e9 + 7)
	for i := 1; i < len(complexity); i++ {
		if complexity[i] <= complexity[0] {
			return 0
		}
		ans = ans * i % mod
	}
	return ans
}
