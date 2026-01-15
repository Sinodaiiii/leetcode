package medium

import "sort"

func maximizeSquareHoleArea260115(n int, m int, hBars []int, vBars []int) int {
	sort.Ints(hBars)
	sort.Ints(vBars)
	hLen := 1
	currLen := 1
	for i := 1; i < len(hBars); i++ {
		if hBars[i] == hBars[i-1]+1 {
			currLen += 1
		} else {
			hLen = max(hLen, currLen)
			currLen = 1
		}
	}
	hLen = max(hLen, currLen)
	vLen := 1
	currLen = 1
	for i := 1; i < len(vBars); i++ {
		if vBars[i] == vBars[i-1]+1 {
			currLen += 1
		} else {
			vLen = max(vLen, currLen)
			currLen = 1
		}
	}
	vLen = max(vLen, currLen)
	// fmt.Println(vLen, hLen)
	width := min(vLen, hLen) + 1
	return width * width
}
