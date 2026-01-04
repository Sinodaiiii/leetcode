package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced260105(root *TreeNode) bool {
	var getDepth func(node *TreeNode) (bool, int)
	getDepth = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}
		lb, ld := getDepth(node.Left)
		rb, rd := getDepth(node.Right)
		return lb && rb && ld-rd <= 1 && ld-rd >= -1, max(ld, rd) + 1
	}

	ret, _ := getDepth(root)
	return ret
}
