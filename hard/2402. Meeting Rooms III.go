package hard

import "sort"

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][0] < meetings[j][0] })
	// fmt.Println(meetings)
	ans := make([]int, n)
	endTime := make([]int, n)
	readyHeap := make([]int, n)
	for i := 0; i < n; i++ {
		readyHeap[i] = i
	}
	busyHeap := make([]int, 0)

	var adjustReadyHeap func(index int)
	var heapUpReadyHeap func(index int)
	var adjustBusyHeap func(index int)
	var heapUpBusyHeap func(index int)

	adjustReadyHeap = func(index int) {
		tmp := index
		if index*2+1 < len(readyHeap) && readyHeap[tmp] > readyHeap[2*index+1] {
			tmp = 2*index + 1
		}
		if index*2+2 < len(readyHeap) && readyHeap[tmp] > readyHeap[2*index+2] {
			tmp = 2*index + 2
		}
		if tmp != index {
			readyHeap[tmp], readyHeap[index] = readyHeap[index], readyHeap[tmp]
			adjustReadyHeap(tmp)
		}
	}
	heapUpReadyHeap = func(index int) {
		for index != 0 {
			father := (index - 1) / 2
			if readyHeap[father] > readyHeap[index] {
				readyHeap[father], readyHeap[index] = readyHeap[index], readyHeap[father]
				index = father
			} else {
				break
			}
		}
	}

	adjustBusyHeap = func(index int) {
		tmp := index
		if index*2+1 < len(busyHeap) && (endTime[busyHeap[tmp]] > endTime[busyHeap[2*index+1]] || (endTime[busyHeap[tmp]] == endTime[busyHeap[2*index+1]] && busyHeap[tmp] > busyHeap[index*2+1])) {
			tmp = 2*index + 1
		}
		if index*2+2 < len(busyHeap) && (endTime[busyHeap[tmp]] > endTime[busyHeap[2*index+2]] || (endTime[busyHeap[tmp]] == endTime[busyHeap[2*index+2]] && busyHeap[tmp] > busyHeap[index*2+2])) {
			tmp = 2*index + 2
		}
		if tmp != index {
			busyHeap[tmp], busyHeap[index] = busyHeap[index], busyHeap[tmp]
			adjustBusyHeap(tmp)
		}
	}
	heapUpBusyHeap = func(index int) {
		for index != 0 {
			father := (index - 1) / 2
			if endTime[busyHeap[father]] > endTime[busyHeap[index]] {
				busyHeap[father], busyHeap[index] = busyHeap[index], busyHeap[father]
				index = father
			} else {
				break
			}
		}
	}

	for _, meeting := range meetings {
		// fmt.Println(meeting)
		for len(busyHeap) > 0 && endTime[busyHeap[0]] <= meeting[0] {
			readyHeap = append(readyHeap, busyHeap[0])
			heapUpReadyHeap(len(readyHeap) - 1)
			busyHeap[0] = busyHeap[len(busyHeap)-1]
			busyHeap = busyHeap[:len(busyHeap)-1]
			adjustBusyHeap(0)
		}
		if len(readyHeap) == 0 {
			ans[busyHeap[0]] += 1
			endTime[busyHeap[0]] += meeting[1] - meeting[0]
			adjustBusyHeap(0)
		} else {
			ans[readyHeap[0]] += 1
			endTime[readyHeap[0]] = max(endTime[readyHeap[0]], meeting[0]) + meeting[1] - meeting[0]
			busyHeap = append(busyHeap, readyHeap[0])
			heapUpBusyHeap(len(busyHeap) - 1)
			readyHeap[0] = readyHeap[len(readyHeap)-1]
			readyHeap = readyHeap[:len(readyHeap)-1]
			adjustReadyHeap(0)
		}
		// fmt.Println("busy: ", busyHeap)
		// fmt.Println("read: ", readyHeap)
		// fmt.Println("end: ", endTime)
		// fmt.Println("ans: ", ans, "\n")
	}
	maxTime := 0
	ret := 0
	for i := 0; i < n; i++ {
		if ans[i] > maxTime {
			maxTime = ans[i]
			ret = i
		}
	}
	return ret
}
