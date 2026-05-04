package medium

import "math"

func minimumDistance(nums []int) int {
	ans := math.MaxInt32
	dict := make([][2]int, len(nums)+1)
	for i, num := range nums {
		if dict[num][1] != 0 {
			ans = min(ans, 2*(i-dict[num][0]+dict[num][1]))
			dict[num] = [2]int{i, i - dict[num][0]}
		} else {
			dict[num] = [2]int{i, math.MaxInt32}
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
