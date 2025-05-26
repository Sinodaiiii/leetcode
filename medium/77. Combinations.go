package medium

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	cur := make([]int, 0)
	var com func(int)
	com = func(num int) {
		if len(cur) == k {
			tmp := make([]int, k)
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}
		if num > n {
			return
		}
		cur = append(cur, num)
		com(num + 1)
		cur = cur[:len(cur)-1]
		com(num + 1)
	}
	com(1)

	return ans
}
