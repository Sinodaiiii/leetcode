package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var getDepth func(root *TreeNode, d int) int
	var depth int
	var ans *TreeNode
	getDepth = func(root *TreeNode, d int) int {
		if root == nil {
			return 0
		}
		ld := getDepth(root.Left, d+1)
		rd := getDepth(root.Right, d+1)
		if ld == rd && d+ld >= depth {
			ans = root
			depth = d + ld
		}
		return max(ld, rd) + 1
	}
	getDepth(root, 0)
	return ans
}
