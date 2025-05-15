package medium

func minOperations(boxes string) []int {
	n := len(boxes)
	left, right := make([]int, n), make([]int, n)
	leftt, rightt := make([]int, n), make([]int, n)
	num, total := 0, 0
	for i := 0; i < n-1; i++ {
		if boxes[i] == '1' {
			num += 1
			total += n - 1 - i
		}
		left[i+1] = num
		rightt[i+1] = total
	}
	num, total = 0, 0
	for i := n - 1; i > 0; i-- {
		if boxes[i] == '1' {
			num += 1
			total += i
		}
		right[i-1] = num
		leftt[i-1] = total
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = leftt[i] - i*right[i] + rightt[i] - (n-1-i)*left[i]
	}
	return ans
}
