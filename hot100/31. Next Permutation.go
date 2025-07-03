package hot100

import "sort"

func nextPermutation(nums []int) {
	n := len(nums)
	if n == 1 {
		return
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			index, nearNum := i+1, nums[i+1]
			for j := i + 1; j < n; j++ {
				if nums[j] > nums[i] && nums[j] < nearNum {
					nearNum = nums[j]
					index = j
				}
			}
			nums[i], nums[index] = nums[index], nums[i]
			sort.Slice(nums[i+1:], func(a, b int) bool { return nums[i+1+a] <= nums[i+1+b] })
			return
		}
	}
	sort.Slice(nums, func(a, b int) bool { return nums[a] <= nums[b] })
}
