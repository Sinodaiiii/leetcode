package graph

import (
	"container/heap"
	"fmt"
	"math"
)

// ==========================================
// 第一部分：基础结构定义
// ==========================================

// edge 代表图中的边
type edge struct {
	to     int // 目标节点
	weight int // 权值（必须非负）
}

// Item 是优先队列中的元素
type Item struct {
	node int // 节点编号
	dist int // 当前从起点到该节点的距离
}

// ==========================================
// 第二部分：实现优先队列 (MinHeap)
// Go 的 heap 需要实现 sort.Interface 加上 Push/Pop
// ==========================================

type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }

// Less 决定了堆的性质。因为我们要找最短路，所以距离小的排前面 (小顶堆)
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// ==========================================
// 第三部分：Dijkstra 核心算法
// ==========================================

// Dijkstra 计算从 start 到图中所有节点的最短路径
// n: 节点总数 (假设节点编号为 0 到 n-1)
// graph: 邻接表，graph[u] 包含了从 u 出发的所有边
// 返回值: dist 数组，dist[i] 表示从 start 到 i 的最短距离
func Dijkstra(n int, start int, graph [][]edge) []int {
	// 1. 初始化距离数组，默认为无穷大
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}

	// 起点到自己的距离为 0
	dist[start] = 0

	// 2. 初始化优先队列
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{node: start, dist: 0})

	// 3. 开始循环
	for pq.Len() > 0 {
		// 取出当前距离最近的节点
		curr := heap.Pop(pq).(Item)
		u := curr.node
		d := curr.dist

		// 【重要优化】懒惰删除
		// 如果从堆中取出的距离 d 大于当前记录的最短距离 dist[u]，
		// 说明这个节点已经被通过更短的路径处理过了，直接跳过。
		if d > dist[u] {
			continue
		}

		// 遍历邻接节点
		for _, edge := range graph[u] {
			v := edge.to
			weight := edge.weight

			// 松弛操作 (Relaxation)
			// 如果 0->u->v 的距离 比 当前记录的 0->v 的距离更短
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				// 将更新后的距离推入堆中
				heap.Push(pq, Item{node: v, dist: dist[v]})
			}
		}
	}

	return dist
}

// ==========================================
// 第四部分：测试运行 (Main)
// ==========================================

func main() {
	// 示例图：5个节点 (0-4)
	// 0 -> 1 (10)
	// 0 -> 4 (5)
	// 1 -> 2 (1)
	// 1 -> 4 (2)
	// 4 -> 1 (3)
	// 4 -> 2 (9)
	// 4 -> 3 (2)
	// 2 -> 3 (4)
	// 3 -> 0 (7)
	// 3 -> 2 (6)

	n := 5
	// 构建邻接表
	graph := make([][]edge, n)

	// 添加边的辅助函数
	addEdge := func(u, v, w int) {
		graph[u] = append(graph[u], edge{to: v, weight: w})
	}

	addEdge(0, 1, 10)
	addEdge(0, 4, 5)
	addEdge(1, 2, 1)
	addEdge(1, 4, 2)
	addEdge(4, 1, 3)
	addEdge(4, 2, 9)
	addEdge(4, 3, 2)
	addEdge(2, 3, 4)
	addEdge(3, 0, 7)
	addEdge(3, 2, 6)

	startNode := 0
	fmt.Printf("正在计算从节点 %d 出发的最短路径...\n", startNode)

	distances := Dijkstra(n, startNode, graph)

	// 打印结果
	for i, d := range distances {
		if d == math.MaxInt {
			fmt.Printf("节点 %d: 不可达\n", i)
		} else {
			fmt.Printf("节点 %d: 最短距离 = %d\n", i, d)
		}
	}
}
