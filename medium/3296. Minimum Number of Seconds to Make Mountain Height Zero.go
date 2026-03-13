package medium

func minNumberOfSeconds260313(mountainHeight int, workerTimes []int) int64 {
	n := len(workerTimes)
	nextTimes := make([]int, n)
	copy(nextTimes, workerTimes)
	total := make([]int, n)
	heap := make([]int, n)
	for i := 0; i < n; i++ {
		heap[i] = i
	}
	adjustHeap := func(index int) {
		for {
			tmp := index
			if index*2+1 < n && (total[heap[index*2+1]]+nextTimes[heap[index*2+1]]) < (total[heap[tmp]]+nextTimes[heap[tmp]]) {
				tmp = index*2 + 1
			}
			if index*2+2 < n && (total[heap[index*2+2]]+nextTimes[heap[index*2+2]]) < (total[heap[tmp]]+nextTimes[heap[tmp]]) {
				tmp = index*2 + 2
			}
			if tmp == index {
				return
			}
			heap[index], heap[tmp] = heap[tmp], heap[index]
			index = tmp
		}
	}
	for i := n/2 - 1; i >= 0; i-- {
		adjustHeap(i)
	}
	// fmt.Println(heap)
	for i := 0; i < mountainHeight; i++ {
		total[heap[0]] += nextTimes[heap[0]]
		nextTimes[heap[0]] += workerTimes[heap[0]]
		// heap[0], heap[n-1] = heap[n-1], heap[0]
		adjustHeap(0)
		// fmt.Println("heap: ", heap)
		// fmt.Println("tota:", total)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, total[i])
	}
	return int64(ans)
}
