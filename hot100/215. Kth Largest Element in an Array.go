package hot100

func findKthLargest(nums []int, k int) int {
	var adjustHeap func(i int)
	adjustHeap = func(i int) {
		tmp := i
		if i*2 < k && nums[i*2] < nums[tmp] {
			tmp = i * 2
		}
		if i*2+1 < k && nums[i*2+1] < nums[tmp] {
			tmp = i*2 + 1
		}
		if tmp != i {
			nums[i], nums[tmp] = nums[tmp], nums[i]
			adjustHeap(tmp)
		}
	}

	for i := k / 2; i >= 0; i-- {
		adjustHeap(i)
	}
	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0] = nums[i]
			adjustHeap(0)
		}
	}
	return nums[0]
}
