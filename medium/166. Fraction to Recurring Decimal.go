package medium

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}
	ans := ""
	if numerator*denominator < 0 {
		ans += "-"
	}
	if numerator < 0 {
		numerator = -numerator
	}
	if denominator < 0 {
		denominator = -denominator
	}
	if numerator > denominator {
		ans += strconv.Itoa(numerator/denominator) + "."
		numerator = numerator % denominator
	} else {
		ans += "0."
	}
	dict := map[[2]int]int{}
	index := len(ans)
	for numerator != 0 {
		numerator *= 10
		if i, exist := dict[[2]int{numerator, denominator}]; exist {
			ans = ans[:i] + "(" + ans[i:] + ")"
			return ans
		}
		dict[[2]int{numerator, denominator}] = index
		ans += strconv.Itoa(numerator / denominator)
		numerator = numerator % denominator
		index += 1
	}

	return ans
}
