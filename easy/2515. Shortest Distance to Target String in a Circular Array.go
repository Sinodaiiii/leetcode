package easy

import "math"

func closestTarget260425(words []string, target string, startIndex int) int {
	if words[startIndex] == target {
		return 0
	}
	n := len(words)
	ans := math.MaxInt32
	for i := startIndex + 1; i <= startIndex+n; i++ {
		// fmt.Println(i, words[i%n])
		if words[i%n] == target {
			ans = min(ans, i-startIndex, n-(i-startIndex))
		}
	}
	if ans < math.MaxInt32 {
		return ans
	}
	return -1
}
