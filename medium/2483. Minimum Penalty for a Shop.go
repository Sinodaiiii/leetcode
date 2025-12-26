package medium

func bestClosingTime(customers string) int {
	n := len(customers)
	totalY := 0
	for i := 0; i < n; i++ {
		if customers[i] == 'Y' {
			totalY += 1
		}
	}
	ans := n
	currMin := n - totalY + 1
	currY := 0
	for i := 0; i < n; i++ {
		// l := i - currY
		// r := totalY - currY
		if totalY+i-2*currY < currMin {
			currMin = totalY + i - 2*currY
			ans = i
		}
		if customers[i] == 'Y' {
			currY += 1
		}
	}
	return ans
}
