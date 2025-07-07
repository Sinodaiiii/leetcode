package hot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var getDepth func(root *TreeNode) int
	getDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := getDepth(root.Left)
		r := getDepth(root.Right)
		ans = max(l+r, ans)
		return max(l, r) + 1
	}

	getDepth(root)
	return ans
}
