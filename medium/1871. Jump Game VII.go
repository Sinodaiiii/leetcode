package medium

func canReach260525(s string, minJump int, maxJump int) bool {
	n := len(s)
	visited := make([]bool, n)
	visited[n-1] = true
	pivot := n - 1
	for i := n - 1; i >= 0; i-- {
		// fmt.Println(i, pivot)
		if visited[i] && s[i] == '0' && pivot >= 0 {
			for pivot > i-minJump && pivot >= 0 {
				pivot -= 1
			}
			for pivot >= i-maxJump && pivot >= 0 {
				visited[pivot] = true
				pivot -= 1
			}
		}
		// fmt.Println(i, pivot, visited)
	}
	return visited[0]
}
