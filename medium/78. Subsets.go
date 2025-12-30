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

func subsets251230(nums []int) [][]int {
	ans := [][]int{[]int{}}
	sort.Ints(nums)
	for _, num := range nums {
		curr := len(ans)
		for i := 0; i < curr; i++ {
			tmpNum := make([]int, len(ans[i]))
			copy(tmpNum, ans[i])
			ans = append(ans, append(tmpNum, num))
		}
	}
	return ans
}
