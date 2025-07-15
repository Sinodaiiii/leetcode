package medium

func countPrimes(n int) int {
	nums := make([]bool, n+1)
	ans := 0
	for i := 2; i < n; i++ {
		if nums[i] == false {
			ans += 1
			for j := 2; i*j < n; j++ {
				nums[i*j] = true
			}
		}
	}
	return ans
}
