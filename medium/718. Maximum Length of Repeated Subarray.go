package medium

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m)
	for _, value := range dp {
		value = make([]int, n)
	}
	if nums1[0] == nums2[0] {
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}
	for i := 1; i < n; i++ {
		if nums1[0] == nums2[i] {
			dp[0][i] = 1
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[0][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {

		}
	}
}
