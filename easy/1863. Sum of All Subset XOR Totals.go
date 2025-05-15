package easy

func subsetXORSum(nums []int) int {
	n := len(nums)
	maxn := 1 << n
	ans := 0
	for i := 0; i < maxn; i++ {
		t := 0
		p := i
		for j := n - 1; j >= 0; j-- {
			t = t ^ ((p & 1) * nums[j])
			p = p >> 1
		}
		ans += t
	}
	return ans
}
