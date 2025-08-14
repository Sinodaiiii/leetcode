package hard

func longestIncreasingPath(matrix [][]int) int {
    ans := 0
    m, n := len(matrix), len(matrix[0])
    long := make([][]int, m)
    for i:=0; i<m; i++ {
        long[i] = make([]int, n)
    }
    var depth func(i, j int) int
    xd := []int{-1, 1, 0, 0}
    yd := []int{0, 0, -1, 1}
    depth = func(x, y int) int {
        ret := 0
        curr := 0
        for index := range xd {
            i, j := x + xd[index], y + yd[index]
            if i >= 0 && i < m && j >= 0 && j < n {
                if matrix[i][j] < matrix[x][y] {
                    if long[i][j] == 0 {
                        curr= depth(i, j)
                    } else {
                        curr = max(curr, long[i][j])
                    }
                    ret = max(ret, curr)
                }
            }
        }
        long[x][y] = ret + 1
        return ret + 1
    }

    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if long[i][j] == 0 {
                ans = max(ans, depth(i, j))
            }
        }
    }
    return ans
}
