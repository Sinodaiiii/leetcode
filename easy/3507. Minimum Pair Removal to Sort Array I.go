package easy

import "math"

func minimumPairRemoval260123(nums []int) int {
	ans := 0
	next := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		next[i] = i + 1
	}
	next[len(next)-1] = -1
	for {
		up := true
		minSum, minIndex := math.MaxInt32, -1
		curr := 0
		for next[curr] != -1 {
			currSum := nums[curr] + nums[next[curr]]
			if currSum < minSum {
				minIndex = curr
				minSum = min(currSum, minSum)
			}
			if nums[curr] > nums[next[curr]] {
				up = false
			}
			curr = next[curr]
		}
		if up {
			return ans
		}
		ans += 1
		nums[minIndex] += nums[next[minIndex]]
		next[minIndex] = next[next[minIndex]]
	}
}
