package medium

func getHappyString260314(n int, k int) string {
	total := 3
	total = total << (n - 1)
	if k > total {
		return ""
	}
	ans, pre := "", 0
	pivot := total / 3
	if k <= pivot {
		ans += "a"
		pre = 0
	} else if k <= 2*pivot {
		ans += "b"
		pre = 1
	} else {
		ans += "c"
		pre = 2
	}
	dict := [][]byte{[]byte{'b', 'c'}, []byte{'a', 'c'}, []byte{'a', 'b'}}
	total = total / 3
	// fmt.Println(total, k, pre, ans)
	for total != 1 {
		k = (k-1)%pivot + 1
		pivot = total / 2
		if k <= pivot {
			ans += string(dict[pre][0])
			pre = int(dict[pre][0] - 'a')
		} else {
			ans += string(dict[pre][1])
			pre = int(dict[pre][1] - 'a')
		}
		// fmt.Println(total, k, pre, ans)
		total = pivot
	}
	return ans
}
