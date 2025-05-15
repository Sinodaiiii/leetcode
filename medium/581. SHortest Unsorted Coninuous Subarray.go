package medium

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	li, ri := 0, 0
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			li = i
			break
		}
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			ri = i
			break
		}
	}
	if li == ri && li == 0 {
		return 0
	}
	minn, maxn := nums[li], nums[ri]
	for i := li; i <= ri; i++ {
		if nums[i] <= minn {
			minn = nums[i]
		}
		if nums[i] >= maxn {
			maxn = nums[i]
		}
	}
	for i := 0; i <= li; i++ {
		if nums[i] > minn {
			li = i
			break
		}
	}
	for i := n - 1; i >= ri; i-- {
		if nums[i] < maxn {
			ri = i
			break
		}
	}
	for i := 0; i <= li; i++ {
		if nums[i] > minn {
			li = i
			break
		}
	}
	return ri - li + 1
}
