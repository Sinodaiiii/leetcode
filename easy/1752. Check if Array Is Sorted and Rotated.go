package easy

func check260523(nums []int) bool {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i%n] > nums[(i+1)%n] {
			for j := i + 1; j < n+i; j++ {
				if nums[j%n] > nums[(j+1)%n] {
					return false
				}
			}
			return true
		}
	}
	return true
}
