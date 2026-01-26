package medium

import "math"

func minCost260127(n int, edges [][]int) int {
	neighbor := make([][][2]int, n)
	distance := make([]int, n)
	for i := 1; i < n; i++ {
		distance[i] = math.MaxInt32
	}
	distance[0] = -1
	for _, edge := range edges {
		neighbor[edge[0]] = append(neighbor[edge[0]], [2]int{edge[1], edge[2]})
		neighbor[edge[1]] = append(neighbor[edge[1]], [2]int{edge[0], edge[2] * 2})
	}
	heap := make([]int, 0, n)

	heapDown := func(index int) {
		i := index
		for {
			tmp := i
			if i*2 < len(heap) && distance[heap[tmp]] > distance[heap[i*2]] {
				tmp = i * 2
			}
			if i*2+1 < len(heap) && distance[heap[tmp]] > distance[heap[i*2+1]] {
				tmp = i*2 + 1
			}
			if tmp == i {
				return
			}
			heap[tmp], heap[i] = heap[i], heap[tmp]
			i = tmp
		}
	}
	heapUp := func(index int) {
		for index > 0 {
			if distance[heap[index]] <= distance[heap[(index-1)/2]] {
				heap[index], heap[(index-1)/2] = heap[(index-1)/2], heap[index]
			} else {
				return
			}
			index = (index - 1) / 2
		}
	}

	curr := 0
	currDist := 0
	for curr != n-1 {
		for _, nei := range neighbor[curr] {
			if distance[nei[0]] > 0 {
				distance[nei[0]] = min(distance[nei[0]], currDist+nei[1])
				heap = append(heap, nei[0])
				heapUp(len(heap) - 1)
			}
		}
		for len(heap) > 0 && distance[heap[0]] < 0 {
			heap[0] = heap[len(heap)-1]
			heap = heap[:len(heap)-1]
			heapDown(0)
		}
		if len(heap) == 0 {
			return -1
		}
		curr = heap[0]
		heap[0] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]
		heapDown(0)
		currDist = distance[curr]
		distance[curr] = -1
	}
	return currDist
}
