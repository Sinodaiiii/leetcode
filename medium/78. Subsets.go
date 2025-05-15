package medium

import "sort"

func subsets(nums []int) [][]int {
	ans := [][]int{}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	n := len(nums)
	ans = append(ans, []int{})
	for i := 0; i < n; i++ {
		lans := len(ans)
		element := nums[i]
		for j := 0; j < lans; j++ {
			newSubset := append([]int{}, ans[j]...)
			newSubset = append(newSubset, element)
			ans = append(ans, newSubset)
		}
	}
	return ans
}
