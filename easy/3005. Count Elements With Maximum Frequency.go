package easy

func maxFrequencyElements(nums []int) int {
	dict := map[int]int{}
	for i := 0; i < len(nums); i++ {
		dict[nums[i]] += 1
	}
	ans := 0
	maxNum := 0
	for _, v := range dict {
		if v > maxNum {
			ans = v
			maxNum = v
		} else if v == maxNum {
			ans += v
		}
	}
	return ans
}
