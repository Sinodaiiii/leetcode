package easy

import "math"

func smallestNumber(n int) int {
	db := n * 2
	num := -1.0
	for db != 0 {
		db /= 2
		num += 1
	}
	return int(math.Pow(2, num)) - 1
}
