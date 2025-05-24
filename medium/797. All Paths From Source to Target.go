func allPathsSourceTarget(graph [][]int) [][]int {
    n := len(graph)
    path := make([]int, 0)
    ans := make([][]int, 0)
    var dfs func(node int)

    dfs = func(node int) {
        path = append(path, node)
        if node == n-1 {
            cur := make([]int, len(path))
            copy(cur, path)
            ans = append(ans, cur)
            path = path[:len(path)-1]
            return
        }
        for _, next := range graph[node] {
            dfs(next)
        }
        path = path[:len(path)-1]
    }

    dfs(0)
    return ans
}