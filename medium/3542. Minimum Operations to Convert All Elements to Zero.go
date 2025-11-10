package medium

func minOperations3542(nums []int) int {
	stack := []int{0}
	left := make([]int, len(nums))
	left[0] = -1
	for i := 1; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] <= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		} else {
			left[i] = -1
		}
		stack = append(stack, i)
	}
	dict := map[int]int{}
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		} else if _, exist := dict[nums[i]]; !exist {
			ans += 1
		} else if dict[nums[i]] < left[i] {
			ans += 1
		}
		dict[nums[i]] = i
	}
	return ans
}
