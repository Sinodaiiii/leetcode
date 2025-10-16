package medium

import "math"

func findSmallestInteger(nums []int, value int) int {
	next := make([]int, value)
	for i := 0; i < value; i++ {
		next[i] = i
	}
	curr := 0
	for _, v := range nums {
		curr = v % value
		if curr < 0 {
			curr += value
		}
		next[curr] += value
	}
	ans := math.MaxInt32
	for _, v := range next {
		ans = min(ans, v)
	}
	return ans
}
