package medium

func longestOnes260213(nums []int, k int) int {
	l, r := 0, 0
	n0 := 0
	ans := 0
	for r < len(nums) {
		if nums[r] == 0 {
			n0 += 1
		}
		if n0 <= k {
			ans = max(ans, r-l+1)
		} else {
			for n0 > k {
				if nums[l] == 0 {
					n0 -= 1
				}
				l += 1
			}
		}
		r += 1
		// fmt.Println(l, r)
	}
	return ans
}
