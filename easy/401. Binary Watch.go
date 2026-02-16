package easy

import "strconv"

func readBinaryWatch260217(turnedOn int) []string {
	countB := func(input int) int {
		ret := 0
		for input != 0 {
			ret += input & 1
			input = input / 2
		}
		return ret
	}
	dictH := map[int][]int{}
	dictM := map[int][]int{}
	for i := 0; i <= 11; i++ {
		index := countB(i)
		dictH[index] = append(dictH[index], i)
	}
	for i := 0; i <= 59; i++ {
		index := countB(i)
		dictM[index] = append(dictM[index], i)
	}
	ans := []string{}
	for i := 0; i <= turnedOn; i++ {
		j := turnedOn - i
		if len(dictH[i]) != 0 && len(dictM[j]) != 0 {
			for _, vh := range dictH[i] {
				h := strconv.Itoa(vh)
				curr := h + ":"
				for _, vm := range dictM[j] {
					m := strconv.Itoa(vm)
					if len(m) == 1 {
						m = "0" + m
					}
					ans = append(ans, curr+m)
				}
			}
		}
	}
	return ans
}
