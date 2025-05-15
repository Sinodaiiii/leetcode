package medium

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	if n == 1 {
		return [][]int{nums}
	}
	ans := make([][]int, 0)
	cur := make([]int, 0)
	dict := map[int]int{}
	keys := make([]int, 0)
	for i := 0; i < n; i++ {
		if dict[nums[i]] == 0 {
			keys = append(keys, nums[i])
		}
		dict[nums[i]]++
	}
	nk := len(keys)
	var backtrack func(l, pre int)
	backtrack = func(level, pre int) {
		if level == n {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}
		for i := 0; i < nk; i++ {
			if i != pre && dict[keys[i]] > 0 {
				tmp := dict[keys[i]]
				tlevel := level
				for dict[keys[i]] != 0 {
					cur = append(cur, keys[i])
					dict[keys[i]]--
					tlevel++
					backtrack(tlevel, i)
				}
				dict[keys[i]] = tmp
				cur = cur[:len(cur)-tmp]
			}
		}
	}
	backtrack(0, -1)
	return ans
}
