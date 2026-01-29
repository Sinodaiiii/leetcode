package graph

import "sort"

// 边结构体
type Edge struct {
	u, v, weight int
}

// 并查集结构
type DSU struct {
	parent []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &DSU{parent: p}
}

func (d *DSU) Find(i int) int {
	if d.parent[i] == i {
		return i
	}
	d.parent[i] = d.Find(d.parent[i]) // 路径压缩
	return d.parent[i]
}

func (d *DSU) Union(i, j int) bool {
	rootI := d.Find(i)
	rootJ := d.Find(j)
	if rootI != rootJ {
		d.parent[rootI] = rootJ
		return true
	}
	return false
}

// Kruskal 算法主函数
func Kruskal(n int, edges []Edge) (int, []Edge) {
	// 1. 按边权升序排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	dsu := NewDSU(n)
	mstWeight := 0
	mstEdges := []Edge{}

	// 2. 遍历边并合并
	for _, edge := range edges {
		if dsu.Union(edge.u, edge.v) {
			mstWeight += edge.weight
			mstEdges = append(mstEdges, edge)
			if len(mstEdges) == n-1 { // 树的边数为 V-1
				break
			}
		}
	}

	if len(mstEdges) < n-1 {
		return -1, nil // 图不连通
	}
	return mstWeight, mstEdges
}
