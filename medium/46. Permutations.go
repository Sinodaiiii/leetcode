package medium

func permute(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{{nums[0]}}
	if n == 1 {
		return ans
	}
	for i := 1; i < n; i++ {
		na := len(ans)
		num := nums[i]
		for j := 0; j < na; j++ {
			nac := len(ans[0])
			tmp := make([]int, nac)
			copy(tmp, ans[j])
			nt := len(tmp)
			for k := 0; k <= nt; k++ {
				ltmp := make([]int, nt+1)
				copy(ltmp[0:k], tmp[0:k])
				copy(ltmp[k+1:nt+1], tmp[k:nt])
				ltmp[k] = num
				ans = append(ans, ltmp)
			}
		}
		ans = ans[na:]
	}
	return ans
}
