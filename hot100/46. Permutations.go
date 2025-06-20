package hot100

func permute(nums []int) [][]int {
	n := len(nums)
	dict := map[int]bool{}
	ans := make([][]int, 0)
	curr := make([]int, 0)
	var dfs func()
	dfs = func() {
		if len(curr) == n {
			tmp := make([]int, n)
			copy(tmp, curr)
			ans = append(ans, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if dict[nums[i]] == false {
				dict[nums[i]] = true
				curr = append(curr, nums[i])
				dfs()
				curr = curr[:len(curr)-1]
				dict[nums[i]] = false
			}
		}
	}

	dfs()
	return ans
}
