package hard

func maxDotProduct260108(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	dp := make([][]int, n1)
	for i := 0; i < n1; i++ {
		dp[i] = make([]int, n2)
	}
	dp[0][0] = nums1[0] * nums2[0]
	for i := 1; i < n2; i++ {
		dp[0][i] = max(dp[0][i-1], nums1[0]*nums2[i])
	}
	for i := 1; i < n1; i++ {
		dp[i][0] = max(dp[i-1][0], nums1[i]*nums2[0])
	}
	for i := 1; i < n1; i++ {
		for j := 1; j < n2; j++ {
			dp[i][j] = max1458(nums1[i]*nums2[j], nums1[i]*nums2[j]+dp[i-1][j-1], dp[i][j-1], dp[i-1][j], dp[i-1][j-1])
		}
	}
	// fmt.Println(dp)
	return dp[n1-1][n2-1]
}

func max1458(nums ...int) int {
	ret := nums[0]
	for _, num := range nums {
		ret = max(ret, num)
	}
	return ret
}
