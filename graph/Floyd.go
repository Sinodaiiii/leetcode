package graph

// Floyd 算法主函数
// matrix 是邻接矩阵，matrix[i][j] 表示 i 到 j 的直接边权，没有边为 math.MaxInt / 2 (防止加法溢出)
func FloydWarshall(n int, matrix [][]int) [][]int {
	// 1. 初始化 dp 数组（这里直接在原矩阵的副本上操作）
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		copy(dist[i], matrix[i])
		dist[i][i] = 0 // 自己到自己的距离为 0
	}

	// 2. 核心动态规划：最外层循环 K 是允许的“中间节点”
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}
