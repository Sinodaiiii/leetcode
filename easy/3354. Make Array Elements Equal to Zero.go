package easy

func countValidSelections(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	curr, ans := 0, 0
	if total%2 == 0 {
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				if curr == total/2 {
					ans += 2
				}
			} else {
				curr += nums[i]
			}
		}
	} else {
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				if curr == total/2 || curr == (total+1)/2 {
					ans += 1
				}
			} else {
				curr += nums[i]
			}
		}
	}
	return ans
}
