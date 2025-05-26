package medium

func removeDuplicates(nums []int) int {
	n := len(nums)
	l, r := 0, 0
	flag := 0
	cur := nums[0]
	for r < n {
		if nums[r] == cur {
			flag += 1
		} else {
			if flag > 2 {
				flag = 2
			}
			for i := 0; i < flag; i++ {
				nums[l] = cur
				l += 1
			}
			cur = nums[r]
			flag = 1
		}
		r += 1
	}
	if flag > 2 {
		flag = 2
	}
	for i := 0; i < flag; i++ {
		nums[l] = cur
		l += 1
	}
	return l
}
