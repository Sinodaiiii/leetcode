package medium

import "math"

func minSumOfLengths260305(arr []int, target int) int {
	preMin := [][2]int{[2]int{-1, math.MaxInt32}}
	index := 0
	l, r, sum := 0, 0, arr[0]
	n := len(arr)
	ans := math.MaxInt32
	for l < n && r < n {
		if sum == target {
			length := r - l + 1
			for index < len(preMin)-1 && l > preMin[index+1][0] {
				index += 1
			}
			ans = min(ans, length+preMin[index][1])
			if length < preMin[len(preMin)-1][1] {
				preMin = append(preMin, [2]int{r, length})
			}
			sum -= arr[l]
			l += 1
		} else if sum < target {
			r += 1
			if r < n {
				sum += arr[r]
			}
		} else {
			sum -= arr[l]
			l += 1
		}
		// fmt.Println(l, r, sum, preMin)
	}
	if ans < math.MaxInt32 {
		return ans
	}
	return -1
}
