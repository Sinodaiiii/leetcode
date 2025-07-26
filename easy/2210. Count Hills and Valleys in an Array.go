package easy

func countHillValley(nums []int) int {
	flag := 0
	ans := 0
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if flag == 0 {
			if nums[i] > pre {
				flag = 1
			} else if nums[i] < pre {
				flag = -1
			}
		} else if nums[i] > pre && flag == -1 {
			flag = 1
			ans += 1
		} else if nums[i] < pre && flag == 1 {
			flag = -1
			ans += 1
		}
		pre = nums[i]
	}
	return ans
}
