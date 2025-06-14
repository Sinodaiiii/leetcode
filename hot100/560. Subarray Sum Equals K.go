package hot100

func subarraySum(nums []int, k int) int {
	pre, ans := 0, 0
	dict := map[int]int{0: 1}
	n := len(nums)
	for i := 0; i < n; i++ {
		pre += nums[i]
		if dict[pre-k] > 0 {
			ans += dict[pre-k]
		}
		dict[pre] += 1
	}
	return ans
}
