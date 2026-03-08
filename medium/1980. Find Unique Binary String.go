package medium

func findDifferentBinaryString260308(nums []string) string {
	dict := map[int]bool{}
	n := len(nums)
	for _, num := range nums {
		curr := 0
		for _, c := range num {
			curr = curr<<1 + int(c-'0')
		}
		dict[curr] = true
	}
	for i := 0; i < (2 << n); i++ {
		if !dict[i] {
			ans := ""
			for i != 0 {
				ans = string(i%2+'0') + ans
				i = i >> 1
			}
			for len(ans) != n {
				ans = "0" + ans
			}
			return ans
		}
	}
	return ""
}
