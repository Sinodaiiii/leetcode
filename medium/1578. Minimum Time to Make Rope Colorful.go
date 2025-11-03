package medium

func minCost(colors string, neededTime []int) int {
	if len(colors) == 0 {
		return 0
	}
	colorBytes := []byte(colors)
	curr := colorBytes[0]
	currMax := neededTime[0]
	ans := 0
	for i := 1; i < len(colorBytes); i++ {
		if colorBytes[i] == curr {
			if neededTime[i] > currMax {
				ans += currMax
				currMax = neededTime[i]
			} else {
				ans += neededTime[i]
			}
		} else {
			curr = colorBytes[i]
			currMax = neededTime[i]
		}
	}
	return ans
}
