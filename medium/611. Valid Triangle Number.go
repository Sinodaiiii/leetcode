package medium

import "sort"

func triangleNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] <= nums[j] })
	ans := 0
	n := len(nums)
	for i := 1; i < n-1; i++ {
		j := 0
		k := i + 1
		for k < n {
			for j < i && nums[j]+nums[i] <= nums[k] {
				j += 1
			}
			if i == j {
				break
			}
			ans += i - j
			// fmt.Println(i, j, k, ans)
			k += 1
		}
	}
	return ans
}
