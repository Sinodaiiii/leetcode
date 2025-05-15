package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countTree(root *TreeNode, distance int) (int, []int) {
	dist := make([]int, distance)
	ret := 0
	if root == nil {
		return ret, dist
	}
	if root.Left == nil && root.Right == nil {
		dist[1] = 1
		return ret, dist
	}
	l, left := countTree(root.Left, distance)
	r, right := countTree(root.Right, distance)
	for i := 1; i < distance; i++ {
		for j := 1; j <= distance-i; j++ {
			ret += left[i] * right[j]
		}
	}
	for i := 2; i < distance; i++ {
		dist[i] = left[i-1] + right[i-1]
	}
	return ret + r + l, dist
}

func countPairs(root *TreeNode, distance int) int {
	if distance == 1 {
		return 0
	}
	ret, _ := countTree(root, distance)
	return ret
}
