package medium

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	//if amount == 0 {
	//	return 0
	//}
	//n := len(coins)
	//sort.Slice(coins, func(i, j int) bool { return coins[i] > coins[j] })
	//if coins[n-1] > amount {
	//	return -1
	//} else if coins[n-1] == amount {
	//	return 1
	//}

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			if i-v >= 0 {
				dp[i] = min(dp[i], dp[i-v]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
