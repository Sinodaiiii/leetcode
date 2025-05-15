package medium

import (
	"math"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n, maxnum := len(nums), int(2*math.Pow(10, 9))
	dict := map[int]int{}
	ans := []int{nums[0]}
	for i := 0; i < n; i++ {
		dict[nums[i]] = 0
	}
	for i := 0; i < n; i++ {
		if count, exist := dict[nums[i]]; exist {
			if count != 0 {
				continue
			}
			curr := []int{}
			t := nums[i]
			dict[t] = 1
			for {
				tmp := t
				curr = append(curr, t)
				for j := 2; t*j < maxnum && t*j <= nums[n-1]; j++ {
					if _, exist2 := dict[t*j]; exist2 {
						dict[t*j] = 1
						t = t * j
						break
					}
				}
				if t == tmp {
					break
				}
			}
			if len(curr) > len(ans) {
				ans = curr
			}
		}
	}
	return ans
}
