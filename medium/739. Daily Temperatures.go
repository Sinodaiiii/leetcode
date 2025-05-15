package medium

func dailyTemperatures(temperatures []int) []int {
	r := len(temperatures) - 1
	l := r - 1
	ans := make([]int, len(temperatures))
	for ; l >= 0; l-- {
		if temperatures[l] >= temperatures[r] {
			ans[l] = 0
			r = l
		} else {
			for i := l + 1; i <= r; i++ {
				if temperatures[i] == temperatures[l] {
					ans[l] = ans[i] + i - l
					break
				} else if temperatures[i] > temperatures[l] {
					ans[l] = i - l
					break
				}
			}
		}
	}
	return ans
}
