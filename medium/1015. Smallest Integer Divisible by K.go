package medium

func smallestRepunitDivByK(k int) int {
	if k%10 == 5 || k%2 == 0 {
		return -1
	}
	ans := 1
	curr := 1
	for curr < k {
		curr = curr*10 + 1
		ans += 1
	}
	rem := curr % k
	for rem != 0 {
		rem = (rem*10 + 1) % k
		if rem < 0 {
			return -1
		} else {
			ans += 1
		}
	}
	return ans
}
