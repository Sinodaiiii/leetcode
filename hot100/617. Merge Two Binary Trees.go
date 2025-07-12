package hot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	v1, v2 := 0, 0
	var l1, r1 *TreeNode
	var l2, r2 *TreeNode
	if root1 != nil {
		v1 = root1.Val
		l1 = root1.Left
		r1 = root1.Right
	}
	if root2 != nil {
		v2 = root2.Val
		l2 = root2.Left
		r2 = root2.Right
	}
	return &TreeNode{v1 + v2, mergeTrees(l1, l2), mergeTrees(r1, r2)}
}
