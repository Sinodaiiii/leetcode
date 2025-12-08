package medium

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum := make([]int, n)
	total := 0
	for i := 0; i < n; i++ {
		sum[i] = gas[i] - cost[i]
		total += sum[i]
	}
	if total < 0 {
		return -1
	}
	if n == 1 {
		return 0
	}
	curr := 0
	total = 0
	for i := 0; i < 2*n; i++ {
		total += sum[i%n]
		if total < 0 {
			curr = i + 1
			total = 0
		}
		if i-curr+2 == n {
			return curr
		}
	}
	return -1
}
