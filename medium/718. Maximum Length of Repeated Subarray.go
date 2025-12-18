package medium

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	// dp := make([][]int, m+1)
	// for i:=0; i<=m; i++ {
	//     dp[i] = make([]int, n+1)
	// }
	dp := make([]int, n+1)
	ans := 0
	for i := 1; i <= m; i++ {
		for j := n; j > 0; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = 0
			}
			ans = max(ans, dp[j])
		}
	}
	return ans
}
