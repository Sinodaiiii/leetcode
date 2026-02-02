package easy

func isTrionic260203(nums []int) bool {
	flag := 0
	direction := true
	if len(nums) <= 3 || nums[1] < nums[0] || nums[len(nums)-1] < nums[len(nums)-2] {
		return false
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return false
		}
		if direction && nums[i] <= nums[i-1] {
			direction = false
			flag += 1
		} else if !direction && nums[i] >= nums[i-1] {
			direction = true
			flag += 1
		}
	}
	return flag == 2
}
