package easy

import (
	"math"
	"sort"
)

func minimumAbsDifference260126(arr []int) [][]int {
	sort.Ints(arr)
	minGap := math.MaxInt32
	for i := 0; i < len(arr)-1; i++ {
		minGap = min(minGap, arr[i+1]-arr[i])
	}
	ans := [][]int{}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == minGap {
			ans = append(ans, []int{arr[i], arr[i+1]})
		}
	}
	return ans
}
