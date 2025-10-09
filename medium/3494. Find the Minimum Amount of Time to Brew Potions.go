package medium

import "math"

func minTime(skill []int, mana []int) int64 {
	n, m := len(skill), len(mana)
	curr := make([]int, n)
	curr[0] = skill[0] * mana[0]
	for i := 1; i < n; i++ {
		curr[i] = curr[i-1] + skill[i]*mana[0]
	}
	for i := 1; i < m; i++ {
		// fmt.Println(curr)
		maxLength := math.MinInt64
		preLength := 0
		maxIndex := -1
		for j := n - 1; j >= 0; j-- {
			preLength += skill[j] * mana[i]
			currLength := curr[j] + preLength
			// fmt.Println("pre: ", preLength)
			// fmt.Println("curr: ", currLength)
			if currLength > maxLength {
				maxLength = currLength
				maxIndex = j
			}
		}
		curr[maxIndex] = curr[maxIndex] + skill[maxIndex]*mana[i]
		for j := maxIndex - 1; j >= 0; j-- {
			curr[j] = curr[j+1] - skill[j+1]*mana[i]
		}
		for j := maxIndex + 1; j < n; j++ {
			curr[j] = curr[j-1] + skill[j]*mana[i]
		}
	}
	return int64(curr[n-1])
}
