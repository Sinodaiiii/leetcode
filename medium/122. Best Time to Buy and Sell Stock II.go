package medium

import "math"

func maxProfit122(prices []int) int {
	ans := 0
	preMax := math.MaxInt32
	preMin := math.MaxInt32
	for _, price := range prices {
		if price < preMax {
			ans += preMax - preMin
			preMin, preMax = price, price
		} else {
			preMax = price
		}
	}
	ans += preMax - preMin
	return ans
}
