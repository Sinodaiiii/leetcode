package hard

func maxCoins(nums []int) int {
	n := len(nums) + 2
	balloon := []int{1}
	balloon = append(balloon, nums...)
	balloon = append(balloon, 1)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	// for i:=2; i<=n-1; i++ {
	//     for j:=0; j<=n-i-1; j++ {
	//         for k:=j+1; k<j+i; k++ {
	//             dp[j][j+i] = max(dp[j][j+i], dp[j][k]+dp[k][j+i]+balloon[j+i]*balloon[k]*balloon[j])
	//         }
	//     }
	// }
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+balloon[i]*balloon[j]*balloon[k])
			}
		}
	}
	return dp[0][n-1]
}
