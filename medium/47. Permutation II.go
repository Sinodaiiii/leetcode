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

func permuteUnique260211(nums []int) [][]int {
	dict := map[int]int{}
	value := make([]int, 0, 8)
	left := make([]int, 0, 8)
	for _, num := range nums {
		if i, exist := dict[num]; exist {
			left[i] += 1
		} else {
			dict[num] = len(value)
			value = append(value, num)
			left = append(left, 1)
		}
	}
	n := len(nums)
	curr := make([]int, 0, 8)
	ans := [][]int{}
	var dfs func()
	dfs = func() {
		if len(curr) == n {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			ans = append(ans, tmp)
			return
		}
		for i, num := range left {
			if num > 0 {
				curr = append(curr, value[i])
				left[i] -= 1
				dfs()
				curr = curr[:len(curr)-1]
				left[i] += 1
			}
		}
	}

	dfs()
	return ans
}
