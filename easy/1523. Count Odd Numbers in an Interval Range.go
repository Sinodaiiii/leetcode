package easy

func countOdds(low int, high int) int {
	if low%2 == 0 {
		return (high+1)/2 - (low+1)/2
	}
	return (high+1)/2 - (low+1)/2 + 1
}
