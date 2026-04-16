package medium

import "math"

func solveQueries260416(nums []int, queries []int) []int {
	dict := map[int]int{}
	n := len(nums)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
	}
	for i := 0; i < n*2; i++ {
		if index, exist := dict[nums[i%n]]; exist && index != i-n {
			dist[i%n] = min(dist[i%n], i-index, index-i+n)
		}
		dict[nums[i%n]] = i
	}
	// fmt.Println(dist)
	dict = map[int]int{}
	for i := 2*n - 1; i >= 0; i-- {
		if index, exist := dict[nums[i%n]]; exist && index != i+n {
			dist[i%n] = min(dist[i%n], index-i, i+n-index)
		}
		dict[nums[i%n]] = i
	}
	// fmt.Println(dist)
	for i, query := range queries {
		if dist[query] == math.MaxInt32 {
			queries[i] = -1
		} else {
			queries[i] = dist[query]
		}
	}
	return queries
}
