package medium

import (
	"fmt"
	"math"
)

func numSub(s string) int {
	m := int(math.Pow(10, 9) + 7)
	fmt.Println(m)
	curr := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			curr += 1
			ans = (ans + curr) % m
		} else {
			curr = 0
		}
	}
	return ans
}
