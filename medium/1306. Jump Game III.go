package medium

func canReach260517(arr []int, start int) bool {
	queue := []int{start}
	arr[start] = -arr[start]
	n := len(arr)
	for len(queue) != 0 {
		tmp := []int{}
		for _, q := range queue {
			offset := -arr[q]
			if q-offset >= 0 && arr[q-offset] >= 0 {
				if arr[q-offset] == 0 {
					return true
				}
				tmp = append(tmp, q-offset)
				arr[q-offset] = -arr[q-offset]
			}
			if q+offset < n && arr[q+offset] >= 0 {
				if arr[q+offset] == 0 {
					return true
				}
				tmp = append(tmp, q+offset)
				arr[q+offset] = -arr[q+offset]
			}
		}
		queue = tmp
	}
	return false
}
