package medium

func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[n-1] = true
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j <= i+nums[i] && j < n; j++ {
			if dp[j] {
				dp[i] = true
			}
		}
	}
	return dp[0]
}

func canJump2(nums []int) bool {
	r := 0
	for curr, num := range nums {
		if curr > r {
			return false
		}
		r = max(r, curr+num)
	}
	return true
}

//func canJump(nums []int) bool {
//	n := len(nums)
//	if n == 1 {
//		return true
//	}
//
//	zero := map[int]bool{}
//	for i := 0; i < n; i++ {
//		if nums[i] == 0 {
//			zero[i] = true
//		}
//	}
//
//	i := 0
//	for j := 0; j < n; j++ {
//		if nums[j] == 0 {
//			if i == j {
//				return false
//			}
//			for zero[i+nums[i]] || i+nums[i] <= j {
//				if i+nums[i] >= n-1 {
//					return true
//				}
//				i++
//				if i == j {
//					return false
//				}
//			}
//			i = j + 1
//		}
//	}
//	return true
//}
