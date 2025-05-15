package medium

import "container/heap"

//func topKFrequent(nums []int, k int) []int {
//	dict := map[int]int{}
//	n := len(nums)
//	for i := 0; i < n; i++ {
//		if _, exist := dict[nums[i]]; exist {
//			dict[nums[i]]++
//		} else {
//			dict[nums[i]] = 1
//		}
//	}
//	i := 0
//	heap := make([]int, 0)
//	for key, value := range dict {
//		i++
//		if i <= k {
//			heap = buildHeap(heap, key, dict)
//		} else {
//			if value > dict[heap[0]] {
//				heap[0] = key
//				adjustHeap(heap, 0, dict)
//			}
//		}
//	}
//	return heap
//}
//
//func buildHeap(heap []int, key int, dict map[int]int) []int {
//	heap = append(heap, key)
//	n := len(heap)
//	for i := (n - 1) / 2; i >= 0; i-- {
//		smallest := i
//		if i*2 < n && dict[heap[i*2]] < dict[heap[smallest]] {
//			smallest = i * 2
//		}
//		if i*2+1 < n && dict[heap[i*2+1]] < dict[heap[smallest]] {
//			smallest = i*2 + 1
//		}
//		if i != smallest {
//			heap[i], heap[smallest] = heap[smallest], heap[i]
//		}
//	}
//	return heap
//}
//
//func adjustHeap(heap []int, index int, dict map[int]int) {
//	n := len(heap)
//	smallest := index
//	if index*2 < n && dict[heap[index*2]] < dict[heap[smallest]] {
//		smallest = index * 2
//	}
//	if index*2+1 < n && dict[heap[index*2+1]] < dict[heap[smallest]] {
//		smallest = index*2 + 1
//	}
//	if index != smallest {
//		heap[index], heap[smallest] = heap[smallest], heap[index]
//		adjustHeap(heap, smallest, dict)
//	}
//}

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
