package graph

import "container/heap"

// 为了简洁，这里重用前面 Dijkstra 问答里的 MinHeap 定义
// type Pair struct { dist, node int }
// type MinHeap []Pair ... (省略重复代码，与 Dijkstra 堆实现一致)

// Prim 算法主函数
func Prim(n int, g [][]Edge) int {
	mstWeight := 0
	visited := make([]bool, n)
	count := 0 // 已加入生成树的节点数

	pq := &MinHeap{}
	heap.Init(pq)
	// 从节点 0 开始，初始权值为 0
	heap.Push(pq, Pair{dist: 0, node: 0})

	for pq.Len() > 0 && count < n {
		curr := heap.Pop(pq).(Pair)
		u, w := curr.node, curr.dist

		if visited[u] {
			continue
		}

		visited[u] = true
		mstWeight += w
		count++

		// 遍历邻接点
		for _, neighbor := range g[u] {
			if !visited[neighbor.to] {
				heap.Push(pq, Pair{dist: neighbor.weight, node: neighbor.to})
			}
		}
	}

	if count < n {
		return -1 // 图不连通
	}
	return mstWeight
}
