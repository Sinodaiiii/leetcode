package medium

func longestConsecutive260105(nums []int) int {
	dict := map[int]bool{}
	n := len(nums)
	for i := 0; i < n; i++ {
		dict[nums[i]] = true
	}
	ans, curr := 0, 0
	for k := range dict {
		curr = 1
		delete(dict, k)
		for j := k - 1; ; j-- {
			if dict[j] {
				curr += 1
				delete(dict, j)
			} else {
				break
			}
		}
		for j := k + 1; ; j++ {
			if dict[j] {
				curr += 1
				delete(dict, j)
			} else {
				break
			}
		}
		ans = max(ans, curr)
	}
	return ans
}
