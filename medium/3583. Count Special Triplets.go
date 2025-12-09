package medium

func specialTriplets(nums []int) int {
	left, right := map[int]int{}, map[int]int{}
	for _, num := range nums {
		right[num] += 1
	}
	ans := 0
	right[nums[0]] -= 1
	for i := 1; i < len(nums)-1; i++ {
		right[nums[i]] -= 1
		left[nums[i-1]] += 1
		ans += left[nums[i]*2] * right[nums[i]*2]
		ans = ans % (1e9 + 7)
	}
	return ans
}
