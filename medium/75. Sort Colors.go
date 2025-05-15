package medium

func sortColors(nums []int) {
	n := len(nums)
	i, j := 0, n-1
	for i != j {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	if nums[i] == 0 {
		i++
	}
	if i == n {
		return
	}
	j = n - 1
	for i != j {
		if nums[i] != 1 {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
}
