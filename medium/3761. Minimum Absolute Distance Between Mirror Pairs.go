package medium

import "math"

func minMirrorPairDistance260418(nums []int) int {
	dict := map[int]int{}
	ans := math.MaxInt32
	for i := len(nums) - 1; i >= 0; i-- {
		re := reverse(nums[i])
		if index, exist := dict[re]; exist {
			ans = min(ans, index-i)
		}
		dict[nums[i]] = i
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func reverse(x int) int {
	ret := 0
	for x != 0 {
		ret = ret*10 + x%10
		x = x / 10
	}
	return ret
}
