package easy

import "fmt"

func hasIncreasingSubarrays(nums []int, k int) bool {
	flag := 1
	for i := 1; i < len(nums); i++ {
		tmpFlag := flag
		if nums[i] > nums[i-1] {
			tmpFlag += 1
		} else {
			tmpFlag = 1
		}
		nums[i-1] = flag
		flag = tmpFlag
	}
	nums[len(nums)-1] = flag
	fmt.Println(nums)
	for i := k - 1; i+k < len(nums); i++ {
		if nums[i] >= k && nums[i+k] >= k {
			return true
		}
	}
	return false
}
