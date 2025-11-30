package medium

import "math"

func minSubarray(nums []int, p int) int {
	ans := math.MaxInt64 / 2
	sum := 0
	for _, num := range nums {
		sum = (sum + num) % p
	}
	target := sum
	if target == 0 {
		return 0
	}
	dict := map[int]int{0: -1}
	sum = 0
	for index, num := range nums {
		sum = (sum + num) % p
		if i, exist := dict[(sum+p-target)%p]; exist {
			ans = min(ans, index-i)
		}
		dict[sum] = index
	}
	if ans == math.MaxInt64/2 || ans == len(nums) {
		return -1
	}
	return ans
}
