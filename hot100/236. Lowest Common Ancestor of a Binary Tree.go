package hot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode = nil
	var findTarget func(root *TreeNode) int
	findTarget = func(root *TreeNode) int {
		if ans != nil || root == nil {
			return 0
		}
		l := findTarget(root.Left)
		r := findTarget(root.Right)
		curr := 0
		if root.Val == p.Val || root.Val == q.Val {
			curr = 1
		}
		if l != 2 && r != 2 && l+r+curr == 2 {
			ans = root
			return 0
		}
		return l + r + curr
	}

	findTarget(root)
	return ans
}
