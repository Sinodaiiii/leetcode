package problems

import "math"

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	f := make([][]int, maxTime+1)
	n := len(passingFees)
	for i := range f {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = math.MaxInt32
		}
	}
	f[0][0] = passingFees[0]
	for t := 1; t <= maxTime; t++ {
		for _, edge := range edges {
			i, j, time := edge[0], edge[1], edge[2]
			if t >= time {
				if f[t-time][i] != math.MaxInt32 {
					f[t][j] = min(f[t][j], f[t-time][i]+passingFees[j])
				}
				if f[t-time][j] != math.MaxInt32 {
					f[t][i] = min(f[t][i], f[t-time][j]+passingFees[i])
				}
			}
		}
	}
	ans := math.MaxInt32
	for i := 1; i <= maxTime; i++ {
		if f[i][n-1] <= ans {
			ans = f[i][n-1]
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
