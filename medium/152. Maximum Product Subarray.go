package medium

func maxProduct(nums []int) int {
	pmax := nums[0]
	nmin := nums[0]
	if nmin >= 0 {
		nmin = 1
	}
	ans := pmax
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		tmp := pmax
		pmax = max(pmax*n, nmin*n, n)
		if pmax > ans {
			ans = pmax
		}
		if nmin*n >= 0 {
			nmin = 1
		} else {
			nmin = min(nmin, tmp*n)
		}
	}

	return ans
}
