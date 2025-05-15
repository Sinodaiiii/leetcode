package medium

func rotate(nums []int, k int) {
	nums = append(nums[len(nums)-k:len(nums)-1], nums[0:len(nums)-k]...)
}
