package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	return isSub(A, B, B)
}

func isSub(A *TreeNode, B *TreeNode, ori *TreeNode) bool {
	if A == nil || B == nil {
		if B == nil {
			return true
		}
		return false
	}
	if A.Val == B.Val {
		l := isSub(A.Left, B.Left, ori)
		r := isSub(A.Right, B.Right, ori)
		if l && r {
			return true
		}
	}
	l := isSub(A.Left, ori, ori)
	r := isSub(A.Right, ori, ori)
	return l || r
}
