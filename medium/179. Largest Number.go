package medium

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		ni, nj := 10, 10
		for nums[i]/ni > 0 {
			ni *= 10
		}
		for nums[j]/nj > 0 {
			nj *= 10
		}
		if nums[i]*nj+nums[j] >= nums[j]*ni+nums[i] {
			return true
		}
		return false
	})
	ans := ""
	for _, num := range nums {
		if ans == "0" {
			if num == 0 {
				continue
			} else {
				ans = ""
			}
		}
		ans += strconv.Itoa(num)
	}
	return ans
}
