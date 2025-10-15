package medium

func maxIncreasingSubarrays(nums []int) int {
	ans := 0
	pre, curr := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			curr += 1
		} else {
			ans = max(ans, min(curr, pre), curr/2)
			pre, curr = curr, 1
		}
	}
	ans = max(ans, min(curr, pre), curr/2)
	return ans
}
