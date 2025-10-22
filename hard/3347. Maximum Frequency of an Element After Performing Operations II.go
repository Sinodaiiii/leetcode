package hard

import "sort"

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	ans := 0
	tail := 0
	n := len(nums)
	for head, num := range nums {
		for tail < n && nums[tail]-k <= num+k {
			tail += 1
		}
		ans = min(max(ans, tail-head), numOperations)
	}
	head := 0
	tail = 0
	same := 0
	for i := 0; i < n; i += same {
		same = 0
		num := nums[i]
		for j := i; j < n && nums[j] == num; j++ {
			same += 1
		}
		for nums[head]+k < num {
			head += 1
		}
		for tail < n && nums[tail]-k <= num {
			tail += 1
		}
		curr := min(max(ans, tail-head), numOperations+same)
		ans = max(curr, ans)
	}
	return ans
}
