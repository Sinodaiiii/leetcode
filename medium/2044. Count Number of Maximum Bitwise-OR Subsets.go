package medium

func countMaxOrSubsets(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	ans := nums[0]
	dict := map[int]int{nums[0]: 1}
	var dfs func(curr, index int)
	dfs = func(curr, index int) {
		if index == len(nums) {
			return
		}
		dfs(curr, index+1)
		or := curr | nums[index]
		dfs(or, index+1)
		dict[or] += 1
		ans = max(ans, or)
	}
	dfs(0, 1)
	dfs(nums[0], 1)
	return dict[ans]
}
