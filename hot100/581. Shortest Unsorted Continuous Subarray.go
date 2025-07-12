package hot100

import "math"

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	nums = append([]int{math.MinInt32}, nums...)
	stack := []int{0}
	l, r := n+1, 0
	ma := nums[0]
	for i := 1; i <= n; i++ {
		if nums[i] > ma {
			ma = nums[i]
		}
		if nums[i] < ma {
			r = i
		}
		if nums[i] < nums[stack[len(stack)-1]] {
			for nums[i] < nums[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			}
			l = min(l, stack[len(stack)-1]+1)
			r = i
			if nums[i] != nums[stack[len(stack)-1]] {
				stack = append(stack, i)
			}
		} else if nums[i] == nums[stack[len(stack)-1]] {
			continue
		} else {
			stack = append(stack, i)
		}
	}
	if l == n+1 {
		return 0
	} else {
		return r - l + 1
	}
}
