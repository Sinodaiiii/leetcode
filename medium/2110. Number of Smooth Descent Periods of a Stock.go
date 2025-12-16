package medium

func getDescentPeriods(prices []int) int64 {
	currL := 1
	ans := 1
	for i := 1; i < len(prices); i++ {
		if prices[i] == prices[i-1]-1 {
			currL += 1
		} else {
			currL = 1
		}
		ans += currL
	}
	return int64(ans)
}
