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

func sortColors2(nums []int) {
	n := len(nums)
	r, b := 0, n-1
	curr := 0
	for curr <= b {
		switch nums[curr] {
		case 0:
			nums[r], nums[curr] = nums[curr], nums[r]
			r += 1
			curr += 1
		case 2:
			nums[b], nums[curr] = nums[curr], nums[b]
			b -= 1
		default:
			curr += 1
		}
	}
}
