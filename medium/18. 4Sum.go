package medium

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var dfs func(nums []int, index, target int)
	n := len(nums)
	ret := [][]int{}
	ans := []int{}
	dict := map[int]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
	anums := make([]int, 0)
	for i := 0; i < n; i++ {
		if dict[nums[i]] == 0 {
			anums = append(anums, nums[i])
		}
		dict[nums[i]]++
	}
	an := len(anums)
	dfs = func(anums []int, index, target int) {
		if target == 0 && len(ans) == 4 {
			tmp := make([]int, 4)
			copy(tmp, ans)
			ret = append(ret, tmp)
			return
		}
		if len(ans) == 4 || index == an {
			return
		}
		dfs(anums, index+1, target)
		l := 4 - len(ans)
		for i := 1; i <= dict[anums[index]] && i <= l; i++ {
			for j := 1; j <= i; j++ {
				ans = append(ans, anums[index])
			}
			dfs(anums, index+1, target-i*anums[index])
			ans = ans[:len(ans)-i]
		}
	}
	dfs(anums, 0, target)
	return ret
}

//func fourSum(nums []int, target int) [][]int {
//	var dfs func(nums []int, index, target int)
//	n := len(nums)
//	ret := [][]int{}
//	ans := []int{}
//	dict := map[string]bool{}
//	sort.Slice(nums, func(i, j int) bool {
//		return nums[i] <= nums[j]
//	})
//	dfs = func(nums []int, index, target int) {
//		if target == 0 && len(ans) == 4 {
//			tmp := make([]int, 4)
//			copy(tmp, ans)
//			key := fmt.Sprint(tmp)
//			if !dict[key] {
//				ret = append(ret, tmp)
//				dict[key] = true
//			}
//			return
//		}
//		if len(ans) == 4 || index == n {
//			return
//		}
//		ans = append(ans, nums[index])
//		dfs(nums, index+1, target-nums[index])
//		ans = ans[:len(ans)-1]
//		dfs(nums, index+1, target)
//	}
//	dfs(nums, 0, target)
//	return ret
//}
